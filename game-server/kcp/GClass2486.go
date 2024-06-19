package kcp

import (
	"encoding/binary"
	"fmt"
	"game-server/helpers"
	"game-server/models"
	"net"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/xtaci/kcp-go/v5"
)

type GClass2486 struct {
	logger          zerolog.Logger
	gclass2487_0    *ConnectionState
	gclass2506_0    *net.UDPConn
	gclass2483_0    *net.UDPAddr
	kcp_0           *kcp.KCP
	gclass2485_0    *GClass2485
	stopwatch_0     time.Time
	gclass2500_0    *GClass2500
	LastReceiveTime uint32
	uint_0          uint32
	uint_1          uint32
	byte_0          []byte
	byte_1          []byte
	ushort_0        uint16
	ushort_1        uint16
	SendQueue       chan *models.GClass2498
	ReceiveQueue    chan *models.GClass2498
	ProfileId       string
	Token           string
	mu              sync.Mutex
}

func NewGClass2486(socket *net.UDPConn, socketAddress *net.UDPAddr, configuration *GClass2485) *GClass2486 {
	g := &GClass2486{
		logger:       zerolog.New(os.Stdout).With().Timestamp().Logger(),
		gclass2506_0: socket,
		gclass2483_0: socketAddress,
		gclass2485_0: configuration,
		stopwatch_0:  time.Now(),
		gclass2500_0: NewGClass2500(),
		byte_0:       make([]byte, 149224),
		byte_1:       make([]byte, 1200),
		ushort_1:     65535,
		SendQueue:    make(chan *models.GClass2498, 100),
		ReceiveQueue: make(chan *models.GClass2498, 16),
	}
	g.gclass2487_0 = NewConnectionState(g)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	g.kcp_0 = kcp.NewKCP(0, g.method_4)
	g.kcp_0.NoDelay(1, 30, 3, 1)
	g.kcp_0.WndSize(256, 256)
	g.kcp_0.SetMtu(1197)

	go g.processReceiveQueue()
	go g.processSendQueue()
	return g
}

func (g *GClass2486) processReceiveQueue() {
	for msg := range g.ReceiveQueue {
		if msg.Type == models.NetworkMessageTypeData {
			HandleDataPacket(msg, g)
		} else {
			fmt.Printf("Unexpected message in receieve queue: %v\n", msg)
		}
	}
}

func (g *GClass2486) processSendQueue() {
	for msg := range g.SendQueue {
		g.SendFinite(msg)
	}
}

func (g *GClass2486) CurrentTime() uint32 {
	return uint32(time.Since(g.stopwatch_0).Milliseconds())
}

func (g *GClass2486) Address() string {
	return g.gclass2483_0.String()
}

func (g *GClass2486) GClass2485_0() *GClass2485 {
	return g.gclass2485_0
}

func (g *GClass2486) Connect() {
	g.logger.Info().Msg(fmt.Sprintf("Connect (address: %s)", g.gclass2483_0))
	g.gclass2487_0.Connect()
}

func (g *GClass2486) Disconnect() {
	g.logger.Info().Msg(fmt.Sprintf("Disconnect (address: %s)", g.gclass2483_0))
	g.gclass2487_0.Disconnect()
}

func (g *GClass2486) EarlyUpdate() {
	g.HandleReceiveReliableFinite()
	g.method_2()
	g.gclass2487_0.Update()
	g.kcp_0.Update()
}

func (g *GClass2486) HandlePingReceiving(buffer []byte, count int) {
	g.method_1(binary.LittleEndian.Uint32(buffer[:4]))
}

func (g *GClass2486) HandlePongReceiving(buffer []byte, count int) {
	num := binary.LittleEndian.Uint32(buffer[:4])
	num2 := g.CurrentTime() - num
	g.gclass2500_0.rtt.Set(float32(num2))
}

func (g *GClass2486) ReturnConnect() {
	g.ReceiveQueue <- g.Get(models.NetworkChannelUnreliable, models.NetworkMessageTypeConnect, nil)
}

func (g *GClass2486) ReturnDisconnect() {
	g.ReceiveQueue <- g.Get(models.NetworkChannelUnreliable, models.NetworkMessageTypeDisconnect, nil)
}

func (g *GClass2486) SendConnect(syn, asc bool) {
	g.logger.Info().Msg(fmt.Sprintf("Send connect (address: %s, syn: %t, asc: %t)", g.gclass2483_0, syn, asc))
	g.SendFinite(g.Get(models.NetworkChannelReliable, models.NetworkMessageTypeConnect, []byte{helpers.BoolToByte(syn), helpers.BoolToByte(asc)}))
}

func (g *GClass2486) SendPing() {
	g.SendFinite(g.Get(models.NetworkChannelUnreliable, models.NetworkMessageTypePing, helpers.UInt32ToBytes(g.CurrentTime())))
}

func (g *GClass2486) SendDisconnect() {
	g.logger.Info().Msg(fmt.Sprintf("Send disconnect (address: %s)", g.gclass2483_0))
	g.SendFinite(g.Get(models.NetworkChannelUnreliable, models.NetworkMessageTypeDisconnect, nil))
}

func (g *GClass2486) HandlePing() {
	if g.CurrentTime()-g.uint_0 > g.gclass2485_0.PingInterval {
		g.SendPing()
		g.uint_0 = g.CurrentTime()
	}
}

func (g *GClass2486) GetInformation() (string, int, byte) {
	ip := g.gclass2483_0
	return ip.IP.String(), ip.Port, 0
}

func (g *GClass2486) GetRtt() int {
	return int(g.gclass2500_0.rtt.averageValue)
}

func (g *GClass2486) GetLossPercent() int {
	return int(g.gclass2500_0.lose.averageLosePercentValue * 100)
}

func (g *GClass2486) GetLossCount() int {
	return int(g.gclass2500_0.lose.averageLoseCountValue)
}

func (g *GClass2486) Send(message *models.GClass2498) {
	g.gclass2487_0.Send(message)
}

func (g *GClass2486) Get(channel models.NetworkChannel, messageType models.NetworkMessageType, buffer []byte) *models.GClass2498 {
	return GetOffset(channel, messageType, buffer, 0, len(buffer))
}

func GetOffset(channel models.NetworkChannel, messageType models.NetworkMessageType, buffer []byte, offset int, count int) *models.GClass2498 {
	array := make([]byte, count)
	copy(array, buffer[offset:offset+count])
	gclass := &models.GClass2498{
		Channel: channel,
		Type:    messageType,
		Buffer:  array[:count],
	}
	return gclass
}

func (g *GClass2486) method_1(t uint32) {
	g.SendFinite(g.Get(models.NetworkChannelUnreliable, models.NetworkMessageTypePong, helpers.UInt32ToBytes(t)))
}

func (g *GClass2486) method_2() {
	if g.CurrentTime()-g.uint_1 > 1000 {
		// g.gclass2500_0.receivedQueue.Set(float32(g.kcp_0.WaitRcv()))
		g.gclass2500_0.sentQueue.Set(float32(g.kcp_0.WaitSnd()))
		g.gclass2500_0.Commit()
		g.uint_1 = g.CurrentTime()
	}
}

func (g *GClass2486) SendFinite(message *models.GClass2498) {
	switch message.Channel {
	case models.NetworkChannelReliable:
		g.method_3(message)
	case models.NetworkChannelUnreliable:
		g.SendUnreliable(message)
	}
}

func (g *GClass2486) SendUnreliable(message *models.GClass2498) {
	num := 4
	num2 := len(g.byte_1) - 4
	count := len(message.Buffer)
	if count > num2 {
		g.logger.Error().Msg(fmt.Sprintf("Unreliable message size to send exceeded [%d/%d] (address: %s)", count, num2, g.gclass2483_0))
		return
	}
	g.ushort_0++
	b := g.byte_1[:num]
	b[0] = byte(message.Channel)
	b[1] = byte(g.ushort_0)
	b[2] = byte(g.ushort_0 >> 8)
	b[3] = byte(message.Type)
	copy(g.byte_1[num:], message.Buffer)
	g.method_5(g.byte_1[:count+num], count+num)
	g.gclass2500_0.unreliableSent.Increment(count)
}

func (g *GClass2486) method_3(message *models.GClass2498) {
	num := 1
	num2 := len(g.byte_0) - 1
	count := len(message.Buffer)
	if count > num2 {
		g.logger.Error().Msg(fmt.Sprintf("Reliable message size to send exceeded [%d/%d] (address: %s)", count, num2, g.gclass2483_0))
		return
	}
	b := g.byte_0[:num+count]
	b[0] = byte(message.Type)
	copy(b[1:], message.Buffer)
	g.kcp_0.Send(b)
	g.gclass2500_0.reliableSent.Increment(count)
}

func (g *GClass2486) method_4(buffer []byte, count int) {
	num := 3
	b := g.byte_1[:num+count]
	g.ushort_0++
	b[0] = 1
	b[1] = byte(g.ushort_0)
	b[2] = byte(g.ushort_0 >> 8)
	copy(b[3:], buffer)
	g.method_5(b[:count+num], count+num)
	g.gclass2500_0.reliableSegmentalSent.Increment(count)
}

func (g *GClass2486) method_5(buffer []byte, count int) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.logger.Debug().Msg(fmt.Sprintf("Writing to UDP (address: %s), data length: %d, data: %x", g.gclass2483_0, count, buffer))
	_, err := g.gclass2506_0.WriteToUDP(buffer[:count], g.gclass2483_0)
	if err != nil {
		g.logger.Warn().Msg(fmt.Sprintf("Error sending data to socket (address: %s): %v", g.gclass2483_0, err))
	} else {
		g.gclass2500_0.sent.Increment(count)
	}
}

func (g *GClass2486) HandleReceive(buffer []byte, count int) {
	g.LastReceiveTime = g.CurrentTime()
	channel := models.NetworkChannel(buffer[0])
	num := binary.LittleEndian.Uint16(buffer[1:3])
	if g.ushort_1 == num {
		g.gclass2500_0.duplicated.Increment()
	} else {
		if g.ushort_1 < num {
			g.ushort_1 = num
		} else {
			if num < g.ushort_1-32767 {
				g.ushort_1 = num
			} else {
				g.gclass2500_0.disordered.Increment()
			}
		}
		g.gclass2500_0.lose.Increment(1, 0)
		if channel == models.NetworkChannelReliable {
			g.method_7(buffer, count)
		} else {
			g.method_6(buffer, count)
		}
		g.gclass2500_0.received.Increment(count)
	}
}

func (g *GClass2486) method_6(buffer []byte, bufferSize int) {
	defer func() {
		if r := recover(); r != nil {
			g.logger.Error().Msg(fmt.Sprintf("Error receiving an unreliable message(address: %s): %v", g.gclass2483_0, r))
		}
	}()
	count := bufferSize - 4
	message := g.Get(models.NetworkChannel(buffer[0]), models.NetworkMessageType(buffer[3]), buffer[4:count+4])
	g.gclass2487_0.HandleReceive(message)
	g.gclass2500_0.unreliableReceived.Increment(bufferSize)
}

func (g *GClass2486) method_7(buffer []byte, bufferSize int) {
	defer func() {
		if r := recover(); r != nil {
			g.logger.Error().Msg(fmt.Sprintf("Error receiving an reliable message(address: %s): %v", g.gclass2483_0, r))
		}
	}()
	num := bufferSize - 3
	// fmt.Printf("bufferSize: %d, num: %d, data: %x\n", bufferSize, num, buffer[3:num+3])
	if num2 := g.kcp_0.Input(buffer[3:num+3], true, true); num2 < 0 {
		g.logger.Error().Msg(fmt.Sprintf("Input failed with error=%d for buffer with length=%d (address: %s)", num2, num, g.gclass2483_0))
		g.logger.Error().Msg(fmt.Sprintf("Data: %x", buffer[:num+3]))

	}
	g.gclass2500_0.reliableSegmentalReceived.Increment(bufferSize)
}

func (g *GClass2486) HandleReceiveReliableFinite() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		g.logger.Error().Msg(fmt.Sprintf("Error finite receiving an reliable message(address: %s): %v", g.gclass2483_0, r))
	// 	}
	// }()
	for {
		if size := g.kcp_0.PeekSize(); size > 0 {
			if n := g.kcp_0.Recv(g.byte_0[:size]); n > 0 {
				messageType := models.NetworkMessageType(g.byte_0[0])
				message := g.Get(models.NetworkChannelReliable, messageType, g.byte_0[1:n])
				g.gclass2487_0.HandleReceive(message)
				g.gclass2500_0.reliableReceived.Increment(n)
			}
		} else {
			break
		}
	}
}

func (g *GClass2486) GetStatistics() GStruct366 {
	return GStruct366{
		Rtt:                       g.gclass2500_0.rtt.averageValue,
		Lose:                      g.gclass2500_0.lose.averageLosePercentValue,
		Disordered:                g.gclass2500_0.disordered.totalValue,
		Duplicated:                g.gclass2500_0.duplicated.totalValue,
		ReliableReceivedAverage:   g.gclass2500_0.reliableReceived.bytes.averageValue,
		ReliableSentAverage:       g.gclass2500_0.reliableSent.bytes.averageValue,
		UnreliableReceivedAverage: g.gclass2500_0.unreliableReceived.bytes.averageValue,
		UnreliableSentAverage:     g.gclass2500_0.unreliableSent.bytes.averageValue,
		ReceivedAverage:           g.gclass2500_0.received.bytes.averageValue,
		SentAverage:               g.gclass2500_0.sent.bytes.averageValue,
		ReceivedTotal:             g.gclass2500_0.received.bytes.totalValue,
		SentTotal:                 g.gclass2500_0.sent.bytes.totalValue,
	}
}

package network

import (
	"encoding/binary"
	"fmt"
	"game-server/helpers"
	"game-server/models"
	"game-server/models/game/enums/network"
	"game-server/network/metrics"
	"net"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/xtaci/kcp-go/v5"
)

type NetworkManager struct {
	logger           zerolog.Logger
	state            *ConnectionState
	conn             *net.UDPConn
	addr             *net.UDPAddr
	kcp_0            *kcp.KCP
	Config           *NetworkConfig
	stopwatch_0      time.Time
	metrics          *metrics.NetworkMetrics
	LastReceiveTime  uint32
	lastPingResponse uint32
	lastUpdateTime   uint32
	byte_0           []byte
	byte_1           []byte
	msgCount         uint16
	lastMsgId        uint16
	SendQueue        chan *models.NetworkMessage
	ReceiveQueue     chan *models.NetworkMessage
	mu               sync.Mutex
}

func NewNetworkManager(socket *net.UDPConn, socketAddress *net.UDPAddr, configuration *NetworkConfig) *NetworkManager {
	g := &NetworkManager{
		logger:       zerolog.New(os.Stdout).With().Timestamp().Logger(),
		conn:         socket,
		addr:         socketAddress,
		Config:       configuration,
		stopwatch_0:  time.Now(),
		metrics:      metrics.NewNetworkMetrics(),
		byte_0:       make([]byte, 149224),
		byte_1:       make([]byte, 1200),
		lastMsgId:    65535,
		SendQueue:    make(chan *models.NetworkMessage, 100),
		ReceiveQueue: make(chan *models.NetworkMessage, 16),
	}
	g.state = NewConnectionState(g)
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	g.kcp_0 = kcp.NewKCP(0, g.outputCallback)
	g.kcp_0.NoDelay(1, 30, 3, 1)
	g.kcp_0.WndSize(256, 256)
	g.kcp_0.SetMtu(1197)

	go g.processReceiveQueue()
	go g.processSendQueue()
	return g
}

func (g *NetworkManager) processReceiveQueue() {
	for msg := range g.ReceiveQueue {
		if msg.Type == models.NetworkMessageTypeData {
			HandleDataPacket(msg, g)
		} else {
			fmt.Printf("Unexpected message in receieve queue: %v\n", msg)
		}
	}
}

func (g *NetworkManager) processSendQueue() {
	for msg := range g.SendQueue {
		g.SendFinite(msg)
	}
}

func (g *NetworkManager) SendReliableDataPacket(hlapi network.PacketID, packet models.GamePacket) {
	resp := &models.DataPacket{
		GamePacketType: uint16(hlapi),
		GamePacket:     packet,
	}
	g.SendQueue <- &models.NetworkMessage{
		Channel: models.NetworkChannelReliable,
		Type:    models.NetworkMessageTypeData,
		Buffer:  resp.Write(),
	}
}

func (g *NetworkManager) CurrentTime() uint32 {
	return uint32(time.Since(g.stopwatch_0).Milliseconds())
}

func (g *NetworkManager) Address() string {
	return g.addr.String()
}

func (g *NetworkManager) Connect() {
	g.logger.Info().Msg(fmt.Sprintf("Connect (address: %s)", g.addr))
	g.state.Connect()
}

func (g *NetworkManager) Disconnect() {
	g.logger.Info().Msg(fmt.Sprintf("Disconnect (address: %s)", g.addr))
	g.state.Disconnect()
}

func (g *NetworkManager) EarlyUpdate() {
	g.HandleReceiveReliableFinite()
	g.updateQueueMetrics()
	g.state.Update()
	g.kcp_0.Update()
}

func (g *NetworkManager) HandlePingReceiving(buffer []byte, count int) {
	g.replyWithPong(binary.LittleEndian.Uint32(buffer[:4]))
}

func (g *NetworkManager) HandlePongReceiving(buffer []byte, count int) {
	num := binary.LittleEndian.Uint32(buffer[:4])
	num2 := g.CurrentTime() - num
	g.metrics.Rtt.Set(float32(num2))
}

func (g *NetworkManager) ReturnConnect() {
	g.ReceiveQueue <- g.Get(models.NetworkChannelUnreliable, models.NetworkMessageTypeConnect, nil)
}

func (g *NetworkManager) ReturnDisconnect() {
	g.ReceiveQueue <- g.Get(models.NetworkChannelUnreliable, models.NetworkMessageTypeDisconnect, nil)
}

func (g *NetworkManager) SendConnect(syn, asc bool) {
	g.logger.Info().Msg(fmt.Sprintf("Send connect (address: %s, syn: %t, asc: %t)", g.addr, syn, asc))
	g.SendFinite(g.Get(models.NetworkChannelReliable, models.NetworkMessageTypeConnect, []byte{helpers.BoolToByte(syn), helpers.BoolToByte(asc)}))
}

func (g *NetworkManager) SendPing() {
	g.SendFinite(g.Get(models.NetworkChannelUnreliable, models.NetworkMessageTypePing, helpers.UInt32ToBytes(g.CurrentTime())))
}

func (g *NetworkManager) SendDisconnect() {
	g.logger.Info().Msg(fmt.Sprintf("Send disconnect (address: %s)", g.addr))
	g.SendFinite(g.Get(models.NetworkChannelUnreliable, models.NetworkMessageTypeDisconnect, nil))
}

func (g *NetworkManager) HandlePing() {
	if g.CurrentTime()-g.lastPingResponse > g.Config.PingInterval {
		g.SendPing()
		g.lastPingResponse = g.CurrentTime()
	}
}

func (g *NetworkManager) GetInformation() (string, int, byte) {
	ip := g.addr
	return ip.IP.String(), ip.Port, 0
}

func (g *NetworkManager) GetRtt() int {
	return int(g.metrics.Rtt.AverageValue)
}

func (g *NetworkManager) GetLossPercent() int {
	return int(g.metrics.Lose.AverageLosePercentValue * 100)
}

func (g *NetworkManager) GetLossCount() int {
	return int(g.metrics.Lose.AverageLoseCountValue)
}

func (g *NetworkManager) Send(message *models.NetworkMessage) {
	g.state.Send(message)
}

func (g *NetworkManager) Get(channel models.NetworkChannel, messageType models.NetworkMessageType, buffer []byte) *models.NetworkMessage {
	return GetOffset(channel, messageType, buffer, 0, len(buffer))
}

func GetOffset(channel models.NetworkChannel, messageType models.NetworkMessageType, buffer []byte, offset int, count int) *models.NetworkMessage {
	array := make([]byte, count)
	copy(array, buffer[offset:offset+count])
	message := &models.NetworkMessage{
		Channel: channel,
		Type:    messageType,
		Buffer:  array[:count],
	}
	return message
}

func (g *NetworkManager) replyWithPong(t uint32) {
	g.SendFinite(g.Get(models.NetworkChannelUnreliable, models.NetworkMessageTypePong, helpers.UInt32ToBytes(t)))
}

func (g *NetworkManager) updateQueueMetrics() {
	if g.CurrentTime()-g.lastUpdateTime > 1000 {
		g.metrics.SentQueue.Set(float32(g.kcp_0.WaitSnd()))
		g.metrics.Commit()
		g.lastUpdateTime = g.CurrentTime()
	}
}

func (g *NetworkManager) SendFinite(message *models.NetworkMessage) {
	switch message.Channel {
	case models.NetworkChannelReliable:
		g.SendReliable(message)
	case models.NetworkChannelUnreliable:
		g.SendUnreliable(message)
	}
}

func (g *NetworkManager) SendUnreliable(message *models.NetworkMessage) {
	num := 4
	num2 := len(g.byte_1) - 4
	count := len(message.Buffer)
	if count > num2 {
		g.logger.Error().Msg(fmt.Sprintf("Unreliable message size to send exceeded [%d/%d] (address: %s)", count, num2, g.addr))
		return
	}
	g.msgCount++
	b := g.byte_1[:num]
	b[0] = byte(message.Channel)
	b[1] = byte(g.msgCount)
	b[2] = byte(g.msgCount >> 8)
	b[3] = byte(message.Type)
	copy(g.byte_1[num:], message.Buffer)
	g.writeToUDP(g.byte_1[:count+num], count+num)
	g.metrics.UnreliableSent.Increment(count)
}

func (g *NetworkManager) SendReliable(message *models.NetworkMessage) {
	num := 1
	num2 := len(g.byte_0) - 1
	count := len(message.Buffer)
	if count > num2 {
		g.logger.Error().Msg(fmt.Sprintf("Reliable message size to send exceeded [%d/%d] (address: %s)", count, num2, g.addr))
		return
	}
	b := g.byte_0[:num+count]
	b[0] = byte(message.Type)
	copy(b[1:], message.Buffer)
	g.kcp_0.Send(b)
	g.metrics.ReliableSent.Increment(count)
}

func (g *NetworkManager) outputCallback(buffer []byte, count int) {
	num := 3
	b := g.byte_1[:num+count]
	g.msgCount++
	b[0] = 1
	b[1] = byte(g.msgCount)
	b[2] = byte(g.msgCount >> 8)
	copy(b[3:], buffer)
	g.writeToUDP(b[:count+num], count+num)
	g.metrics.ReliableSegmentalSent.Increment(count)
}

func (g *NetworkManager) writeToUDP(buffer []byte, count int) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.logger.Debug().Msg(fmt.Sprintf("Writing to UDP (address: %s), data length: %d, data: %x", g.addr, count, buffer))
	_, err := g.conn.WriteToUDP(buffer[:count], g.addr)
	if err != nil {
		g.logger.Warn().Msg(fmt.Sprintf("Error sending data to socket (address: %s): %v", g.addr, err))
	} else {
		g.metrics.Sent.Increment(count)
	}
}

func (g *NetworkManager) HandleReceive(buffer []byte, count int) {
	g.LastReceiveTime = g.CurrentTime()
	channel := models.NetworkChannel(buffer[0])
	num := binary.LittleEndian.Uint16(buffer[1:3])
	if g.lastMsgId == num {
		g.metrics.Duplicated.Increment()
	} else {
		if g.lastMsgId < num {
			g.lastMsgId = num
		} else {
			if num < g.lastMsgId-32767 {
				g.lastMsgId = num
			} else {
				g.metrics.Disordered.Increment()
			}
		}
		g.metrics.Lose.Increment(1, 0)
		if channel == models.NetworkChannelReliable {
			g.ReceiveReliable(buffer, count)
		} else {
			g.ReceiveUnreliable(buffer, count)
		}
		g.metrics.Received.Increment(count)
	}
}

func (g *NetworkManager) ReceiveUnreliable(buffer []byte, bufferSize int) {
	defer func() {
		if r := recover(); r != nil {
			g.logger.Error().Msg(fmt.Sprintf("Error receiving an unreliable message(address: %s): %v", g.addr, r))
		}
	}()
	count := bufferSize - 4
	message := g.Get(models.NetworkChannel(buffer[0]), models.NetworkMessageType(buffer[3]), buffer[4:count+4])
	g.state.HandleReceive(message)
	g.metrics.UnreliableReceived.Increment(bufferSize)
}

func (g *NetworkManager) ReceiveReliable(buffer []byte, bufferSize int) {
	defer func() {
		if r := recover(); r != nil {
			g.logger.Error().Msg(fmt.Sprintf("Error receiving an reliable message(address: %s): %v", g.addr, r))
		}
	}()
	num := bufferSize - 3
	if num2 := g.kcp_0.Input(buffer[3:num+3], true, true); num2 < 0 {
		g.logger.Error().Msg(fmt.Sprintf("Input failed with error=%d for buffer with length=%d (address: %s)", num2, num, g.addr))
		g.logger.Error().Msg(fmt.Sprintf("Data: %x", buffer[:num+3]))

	}
	g.metrics.ReliableSegmentalReceived.Increment(bufferSize)
}

func (g *NetworkManager) HandleReceiveReliableFinite() {
	defer func() {
		if r := recover(); r != nil {
			g.logger.Error().Msg(fmt.Sprintf("Error finite receiving an reliable message(address: %s): %v", g.addr, r))
		}
	}()
	for {
		if size := g.kcp_0.PeekSize(); size > 0 {
			if n := g.kcp_0.Recv(g.byte_0[:size]); n > 0 {
				messageType := models.NetworkMessageType(g.byte_0[0])
				message := g.Get(models.NetworkChannelReliable, messageType, g.byte_0[1:n])
				g.state.HandleReceive(message)
				g.metrics.ReliableReceived.Increment(n)
			}
		} else {
			break
		}
	}
}

package kcp

import (
	"fmt"

	"github.com/rs/zerolog/log"
)

type Logger struct{}

func (l *Logger) LogInfo(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
}

func (l *Logger) LogError(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
}

type Message struct {
	Type    NetworkMessageType
	Buffer  []byte
	Dispose func()
}

type State int

const (
	Initial State = iota
	Connecting
	Connected
	Waiting
	Disconnected
)

type ConnectionState struct {
	Connection *GClass2486
	State      State
	StartTime  float64
}

func NewConnectionState(conn *GClass2486) *ConnectionState {
	return &ConnectionState{Connection: conn}
}

func (cs *ConnectionState) Enter(state State) {
	cs.State = state
	switch state {
	case Initial:
		log.Info().Msg(fmt.Sprintf("Enter to the 'Initial' state (address: %s)", cs.Connection.Address()))
	case Connecting:
		log.Info().Msg(fmt.Sprintf("Enter to the 'Connecting' state (address: %s)", cs.Connection.Address()))
		// cs.StartTime = cs.Connection.CurrentTime
	case Connected:
		log.Info().Msg(fmt.Sprintf("Enter to the 'Connected' state (address: %s)", cs.Connection.Address()))
		// cs.Connection.LastReceiveTime = cs.Connection.CurrentTime
	case Waiting:
		log.Info().Msg(fmt.Sprintf("Enter to the 'Waiting' state (address: %s)", cs.Connection.Address()))
	}
}

func (cs *ConnectionState) Exit() {
	log.Info().Msg(fmt.Sprintf("Exit state (address: %s)", cs.Connection.Address()))
}

func (cs *ConnectionState) HandleReceive(msg *GClass2498) {
	switch cs.State {
	case Initial:
		if msg.Type == NetworkMessageTypeConnect {
			cs.Enter(Connecting)
			cs.Connection.SendConnect(true, true)
			// msg.Dispose()
		}
	case Connecting:
		if msg.Type == NetworkMessageTypeConnect {
			cs.Enter(Connected)
			if msg.Buffer[0] == 1 {
				cs.Connection.SendConnect(false, true)
			}
			cs.Connection.ReturnConnect()
			// msg.Dispose()
		} else if msg.Type == NetworkMessageTypeData {
			cs.Connection.ReceiveQueue <- msg
		}
		// else {
		// msg.Dispose()
		// }
	case Connected:
		switch msg.Type {
		case NetworkMessageTypePing:
			cs.Connection.HandlePingReceiving(msg.Buffer, len(msg.Buffer))
			// msg.Dispose()
		case NetworkMessageTypePong:
			cs.Connection.HandlePongReceiving(msg.Buffer, len(msg.Buffer))
			// msg.Dispose()
		case NetworkMessageTypeData:
			cs.Connection.ReceiveQueue <- msg
		case NetworkMessageTypeDisconnect:
			log.Info().Msg(fmt.Sprintf("Receive disconnect (address: %s)", cs.Connection.Address()))
			cs.Connection.ReturnDisconnect()
			cs.Enter(Disconnected)
			// msg.Dispose()
		default:
			// msg.Dispose()
		}
	case Waiting:
		if msg.Type == NetworkMessageTypeConnect {
			cs.Enter(Connected)
			cs.HandleReceive(msg)
		}
	}
}

func (cs *ConnectionState) Update() {
	switch cs.State {
	case Connecting, Waiting:
		cs.HandleTimeout()
	}
	if cs.State == Connected {
		cs.Connection.HandlePing()
		// cs.Connection.HandleDeadLink()
		// cs.Connection.HandleOverflow()
		cs.FlushSendQueue()
	}
}

func (cs *ConnectionState) Connect() {
	cs.Enter(Connecting)
	cs.Connection.SendConnect(true, false)
}

func (cs *ConnectionState) Disconnect() {
	cs.Connection.SendDisconnect()
	cs.Connection.ReturnDisconnect()
	cs.Enter(Disconnected)
}

func (cs *ConnectionState) Send(msg *GClass2498) {
	cs.Connection.SendQueue <- msg
}

func (cs *ConnectionState) HandleTimeout() {
	// if time.Since(time.Unix(int64(cs.Connection.LastReceiveTime), 0)).Seconds() > 5 {
	// 	log.Error().Msg(fmt.Sprintf("Timeout: Connection timed out after not receiving any message (address: %s)", cs.Connection.Address))
	// 	cs.Disconnect()
	// }
}

func (cs *ConnectionState) FlushSendQueue() {
	for len(cs.Connection.SendQueue) > 0 {
		msg := <-cs.Connection.SendQueue
		cs.Connection.SendFinite(msg)
	}
}

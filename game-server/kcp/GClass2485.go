package kcp

type GClass2485 struct {
	ConnectionLimit   int
	PacketSize        int
	ConnectingTimeout uint
	WaitTimeout       uint
	DisconnectTimeout uint
	PingInterval      uint32
	SendWindowSize    uint
	ReceiveWindowSize uint
	SendPoolLimit     int
	ReceivePoolLimit  int
}

const (
	CONNECTION_LIMIT    = 32
	PACKET_SIZE         = 2500
	CONNECTING_TIMEOUT  = 4000
	WAIT_TIMEOUT        = 3000
	DISCONNECT_TIMEOUT  = 12000
	PING_INTERVAL       = 1000
	SEND_WINDOW_SIZE    = 256
	RECEIVE_WINDOW_SIZE = 256
	SEND_POOL_LIMIT     = 8192
	RECEIVE_POOL_LIMIT  = 8192
)

func NewGClass2485() *GClass2485 {
	return &GClass2485{
		ConnectionLimit:   CONNECTION_LIMIT,
		PacketSize:        PACKET_SIZE,
		ConnectingTimeout: CONNECTING_TIMEOUT,
		WaitTimeout:       WAIT_TIMEOUT,
		DisconnectTimeout: DISCONNECT_TIMEOUT,
		PingInterval:      PING_INTERVAL,
		SendWindowSize:    SEND_WINDOW_SIZE,
		ReceiveWindowSize: RECEIVE_WINDOW_SIZE,
		SendPoolLimit:     SEND_POOL_LIMIT,
		ReceivePoolLimit:  RECEIVE_POOL_LIMIT,
	}
}

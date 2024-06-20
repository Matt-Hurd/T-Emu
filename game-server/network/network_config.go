package network

type NetworkConfig struct {
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

func DefaultNetworkConfig() *NetworkConfig {
	return &NetworkConfig{
		ConnectionLimit:   32,
		PacketSize:        2500,
		ConnectingTimeout: 4000,
		WaitTimeout:       3000,
		DisconnectTimeout: 12000,
		PingInterval:      1000,
		SendWindowSize:    256,
		ReceiveWindowSize: 256,
		SendPoolLimit:     8192,
		ReceivePoolLimit:  8192,
	}
}

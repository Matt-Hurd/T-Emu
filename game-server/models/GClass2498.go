package models

type NetworkChannel byte

const (
	NetworkChannelNone NetworkChannel = iota
	NetworkChannelReliable
	NetworkChannelUnreliable
)

type NetworkMessageType byte

const (
	NetworkMessageTypeNone NetworkMessageType = iota
	NetworkMessageTypeConnect
	NetworkMessageTypePing
	NetworkMessageTypePong
	NetworkMessageTypeData
	NetworkMessageTypeDisconnect
)

type GClass2498 struct {
	Channel NetworkChannel
	Type    NetworkMessageType
	Buffer  []byte
}

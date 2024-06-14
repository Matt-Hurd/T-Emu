package handlers

import (
	"net"
)

type PacketHandler interface {
	HandlePacket(data []byte, addr *net.UDPAddr)
}

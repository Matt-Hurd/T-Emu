package handlers

import (
	"log"
	"net"
)

type UnreliableHandler struct{}

func NewUnreliableHandler() *UnreliableHandler {
	return &UnreliableHandler{}
}

func (h *UnreliableHandler) HandlePacket(data []byte, addr *net.UDPAddr) {
	log.Printf("Received Unreliable packet from %s: %x\n", addr, data)
}

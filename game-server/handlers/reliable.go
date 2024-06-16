package handlers

import (
	"net"

	"game-server/kcp"
)

type ReliableHandler struct {
	kcpHandler *kcp.KCPHandler
}

func NewReliableHandler(kcpHandler *kcp.KCPHandler) *ReliableHandler {
	return &ReliableHandler{kcpHandler: kcpHandler}
}

func (h *ReliableHandler) HandlePacket(data []byte, addr *net.UDPAddr) {
	h.kcpHandler.HandleReliablePacket(data, addr)
}

package kcp

import (
	"log"
	"net"
	"sync"

	"game-server/models"
	"game-server/udp"

	kcp "github.com/xtaci/kcp-go/v5"
)

type KCPHandler struct {
	kcpSessions map[string]*kcp.UDPSession
	udpConns    map[string]*udp.UDPConn
	mutex       sync.Mutex
	conn        *net.UDPConn
}

func NewKCPHandler(conn *net.UDPConn) *KCPHandler {
	return &KCPHandler{
		kcpSessions: make(map[string]*kcp.UDPSession),
		udpConns:    make(map[string]*udp.UDPConn),
		conn:        conn,
	}
}

func (h *KCPHandler) HandlePacket(data []byte, addr *net.UDPAddr) {
	if len(data) == 0 {
		return
	}

	switch data[0] {
	case 1:
		h.handleReliablePacket(data[3:], addr)
	case 2:
		h.handleUnreliablePacket(data[3:], addr)
	}
}

func (h *KCPHandler) handleReliablePacket(data []byte, addr *net.UDPAddr) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	udpConn, exists := h.udpConns[addr.String()]
	if !exists {
		udpConn = &udp.UDPConn{
			Conn:       h.conn,
			Buffer:     make([]byte, len(data)),
			Addr:       addr,
			ReadSignal: make(chan struct{}, 1),
		}
		copy(udpConn.Buffer, data)
		h.udpConns[addr.String()] = udpConn

		var err error
		session, err := kcp.NewConn3(0, addr, nil, 10, 3, udpConn)
		if err != nil {
			log.Printf("Error creating KCP connection: %v", err)
			return
		}
		session.SetNoDelay(1, 20, 2, 1)
		session.SetWindowSize(256, 256)
		session.SetMtu(1197)
		h.kcpSessions[addr.String()] = session
		go h.handleKCPSession(session)
	} else {
		udpConn.Mutex.Lock()
		udpConn.Buffer = make([]byte, len(data))
		copy(udpConn.Buffer, data)
		udpConn.Addr = addr
		udpConn.Mutex.Unlock()
	}
	udpConn.ReadSignal <- struct{}{}
}

func (h *KCPHandler) handleUnreliablePacket(data []byte, addr *net.UDPAddr) {
	log.Printf("Received Unreliable packet from %s: %x\n", addr, data)
}

func (h *KCPHandler) handleKCPSession(session *kcp.UDPSession) {
	defer session.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := session.Read(buffer)
		if err != nil {
			log.Printf("Error reading from KCP connection: %v", err)
			return
		}

		packet, err := models.ParsePacket(buffer[:n])
		if err != nil {
			log.Printf("Error parsing packet: %v", err)
			continue
		}

		switch p := packet.(type) {
		case *models.ConnectPacket:
			h.handleConnectPacket(session, p)
		default:
			log.Printf("Unhandled packet type: %d", packet.Type())
		}
	}
}

func (h *KCPHandler) handleConnectPacket(session *kcp.UDPSession, packet *models.ConnectPacket) {
	log.Printf("Received ConnectPacket: Syn=%x, Asc=%x", packet.Syn, packet.Asc)
	session.Write([]byte{1, 1})
	// Handle login logic here
}

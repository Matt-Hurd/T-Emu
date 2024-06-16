package kcp

import (
	"log"
	"net"
	"sync"
	"time"

	"game-server/controllers"
	"game-server/udp"

	kcp "github.com/xtaci/kcp-go/v5"
)

type KCPHandler struct {
	KcpSessions map[string]*kcp.UDPSession
	UdpConns    map[string]*udp.UDPConn
	mutex       sync.Mutex
	conn        *net.UDPConn
}

func NewKCPHandler(conn *net.UDPConn) *KCPHandler {
	return &KCPHandler{
		KcpSessions: make(map[string]*kcp.UDPSession),
		UdpConns:    make(map[string]*udp.UDPConn),
		conn:        conn,
	}
}

func (h *KCPHandler) HandleReliablePacket(data []byte, addr *net.UDPAddr) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	udpConn, exists := h.UdpConns[addr.String()]
	if !exists {
		udpConn = &udp.UDPConn{
			Conn:       h.conn,
			Buffer:     make([]byte, len(data)),
			Addr:       addr,
			ReadSignal: make(chan struct{}, 1),
			Start:      time.Now(),
		}
		copy(udpConn.Buffer, data)
		h.UdpConns[addr.String()] = udpConn

		var err error
		session, err := kcp.NewConn3(0, addr, nil, 10, 3, udpConn)
		if err != nil {
			log.Printf("Error creating KCP connection: %v", err)
			return
		}
		session.SetNoDelay(1, 20, 2, 1)
		session.SetWindowSize(256, 256)
		session.SetMtu(1205)
		h.KcpSessions[addr.String()] = session
		go h.handleKCPSession(session, udpConn)
	} else {
		udpConn.Mutex.Lock()
		udpConn.Buffer = make([]byte, len(data))
		copy(udpConn.Buffer, data)
		udpConn.Addr = addr
		udpConn.Mutex.Unlock()
	}
	udpConn.ReadSignal <- struct{}{}
}

func (h *KCPHandler) handleKCPSession(session *kcp.UDPSession, conn *udp.UDPConn) {
	defer session.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := session.Read(buffer)
		if err != nil {
			log.Printf("Error reading from KCP connection: %v", err)
			return
		}

		controllers.HandlePacket(session, buffer[:n], conn)
	}
}

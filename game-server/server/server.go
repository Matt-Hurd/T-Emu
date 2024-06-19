package server

import (
	"log"
	"net"
	"time"

	"game-server/kcp"
)

type Server struct {
	// kcpHandler      *kcp.KCPHandler
	// reliableHandler *handlers.ReliableHandler
	clients map[string]*kcp.GClass2486
	conn    *net.UDPConn
}

func NewServer() *Server {
	return &Server{clients: make(map[string]*kcp.GClass2486)}
}

func (srv *Server) Start() error {
	addr := net.UDPAddr{
		Port: 9090,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		return err
	}
	srv.conn = conn
	defer conn.Close()

	// srv.kcpHandler = kcp.NewKCPHandler(conn)
	// srv.reliableHandler = handlers.NewReliableHandler(srv.kcpHandler)

	log.Println("Combined KCP/UDP server listening on port 9090")

	go srv.PacketHandlerLoop()
	for {
		time.Sleep(10 * time.Millisecond)
		for _, conn := range srv.clients {
			conn.EarlyUpdate()
		}
	}
}

func (srv *Server) PacketHandlerLoop() {
	buffer := make([]byte, 65535)
	for {
		n, remoteAddr, err := srv.conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error reading from UDP connection: %v", err)
			continue
		}

		go srv.handlePacket(buffer[:n], n, remoteAddr)
	}
}

func (srv *Server) handlePacket(data []byte, n int, addr *net.UDPAddr) {
	if len(data) == 0 {
		return
	}

	// fmt.Printf("Received packet from %s: %x\n", addr, data)
	// fmt.Printf("Received packet from %s: %x\n", addr.String(), data[:n])
	if _, exists := srv.clients[addr.String()]; !exists {
		srv.clients[addr.String()] = kcp.NewGClass2486(srv.conn, addr, kcp.NewGClass2485())
	}
	srv.clients[addr.String()].HandleReceive(data, n)
	srv.clients[addr.String()].HandleReceiveReliableFinite()

	// switch data[0] {
	// case 1:
	// 	srv.reliableHandler.HandlePacket(data, addr)
	// case 2:
	// 	controllers.HandlePacket(srv.kcpHandler.KcpSessions[addr.String()], data, srv.kcpHandler.UdpConns[addr.String()])
	// default:
	// 	log.Printf("Unknown packet type: %d\n", data[0])
	// }
}

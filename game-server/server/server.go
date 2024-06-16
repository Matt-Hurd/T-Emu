package server

import (
	"fmt"
	"log"
	"net"
	"time"

	"game-server/controllers"
	"game-server/handlers"
	"game-server/kcp"
	"game-server/models"
)

type Server struct {
	kcpHandler      *kcp.KCPHandler
	reliableHandler *handlers.ReliableHandler
	conn            *net.UDPConn
}

func NewServer() *Server {
	return &Server{}
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

	srv.kcpHandler = kcp.NewKCPHandler(conn)
	srv.reliableHandler = handlers.NewReliableHandler(srv.kcpHandler)

	log.Println("Combined KCP/UDP server listening on port 9090")

	go srv.PacketHandlerLoop()
	for {
		time.Sleep(500 * time.Millisecond)
		for _, conn := range srv.kcpHandler.UdpConns {
			packet := &models.PingPacket{}
			conn.WriteToUnreliable(packet.Write(uint32(time.Since(conn.Start).Milliseconds())), conn.Addr)
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

		go srv.handlePacket(buffer[:n], remoteAddr)
	}
}

func (srv *Server) handlePacket(data []byte, addr *net.UDPAddr) {
	if len(data) == 0 {
		return
	}

	fmt.Printf("Received packet from %s: %x\n", addr, data)

	switch data[0] {
	case 1:
		srv.reliableHandler.HandlePacket(data, addr)
	case 2:
		controllers.HandlePacket(srv.kcpHandler.KcpSessions[addr.String()], data, srv.kcpHandler.UdpConns[addr.String()])
	default:
		log.Printf("Unknown packet type: %d\n", data[0])
	}
}

package server

import (
	"log"
	"net"
	"time"

	"game-server/models/player"
	"game-server/network"
)

type Server struct {
	clients map[string]*player.PlayerSession
	conn    *net.UDPConn
}

func NewServer() *Server {
	return &Server{clients: make(map[string]*player.PlayerSession)}
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

	log.Println("Combined KCP/UDP server listening on port 9090")

	go srv.PacketHandlerLoop()
	for {
		time.Sleep(10 * time.Millisecond)
		for _, conn := range srv.clients {
			conn.Connection.EarlyUpdate()
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

	if _, exists := srv.clients[addr.String()]; !exists {
		srv.clients[addr.String()] = &player.PlayerSession{
			Connection: network.NewNetworkManager(srv.conn, addr, network.DefaultNetworkConfig()),
		}
	}
	srv.clients[addr.String()].Connection.HandleReceive(data, n)

}

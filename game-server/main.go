package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	kcp "github.com/xtaci/kcp-go/v5"
)

type udpConn struct {
	conn       *net.UDPConn
	buffer     []byte
	addr       *net.UDPAddr
	mutex      sync.Mutex
	readSignal chan struct{}
}

func (c *udpConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	// Wait for data to be available
	<-c.readSignal

	c.mutex.Lock()
	defer c.mutex.Unlock()

	if len(c.buffer) > 0 {
		n = copy(p, c.buffer)
		addr = c.addr
		fmt.Printf("Received %d bytes (hack): %x\n", n, p[:n])
		c.buffer = nil // Clear the buffer once it has been read
		return n, addr, nil
	}
	fmt.Println("no data in buffer")
	return 0, nil, nil // No data in buffer
}

func (c *udpConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	fmt.Println("WriteTo")
	return c.conn.WriteTo(p, addr)
}

func (c *udpConn) Close() error {
	fmt.Println("Close")
	return c.conn.Close()
}

func (c *udpConn) LocalAddr() net.Addr {
	fmt.Println("LocalAddr")
	return c.conn.LocalAddr()
}

func (c *udpConn) SetDeadline(t time.Time) error {
	fmt.Println("SetDeadline")
	return c.conn.SetDeadline(t)
}

func (c *udpConn) SetReadDeadline(t time.Time) error {
	fmt.Println("SetReadDeadline")
	return c.conn.SetReadDeadline(t)
}

func (c *udpConn) SetWriteDeadline(t time.Time) error {
	fmt.Println("SetWriteDeadline")
	return c.conn.SetWriteDeadline(t)
}

type server struct {
	kcpSessions map[string]*kcp.UDPSession
	udpConns    map[string]*udpConn
	mutex       sync.Mutex
	conn        *net.UDPConn
}

func main() {
	srv := &server{
		kcpSessions: make(map[string]*kcp.UDPSession),
		udpConns:    make(map[string]*udpConn),
	}
	srv.start()
}

func (srv *server) start() {
	addr := net.UDPAddr{
		Port: 9090,
		IP:   net.ParseIP("0.0.0.0"),
	}
	conn, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatalf("Error setting up UDP listener: %v", err)
	}
	srv.conn = conn
	defer conn.Close()

	log.Println("Combined KCP/UDP server listening on port 9090")

	buffer := make([]byte, 65535)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			log.Printf("Error reading from UDP connection: %v", err)
			continue
		}

		go srv.handlePacket(buffer[:n], remoteAddr)
	}
}

func (srv *server) handlePacket(data []byte, addr *net.UDPAddr) {
	if len(data) == 0 {
		return
	}

	// switch data[0]{
	// case 1:
	srv.handleReliablePacket(data[3:], addr)
	// case 2:
	// 	srv.handleUnreliablePacket(data, addr)
	// }
}

func (srv *server) handleReliablePacket(data []byte, addr *net.UDPAddr) {
	srv.mutex.Lock()
	defer srv.mutex.Unlock()

	udpConnection, exists := srv.udpConns[addr.String()]
	if !exists {
		log.Printf("Creating new KCP session for %s\n", addr)
		udpConnection = &udpConn{
			conn:       srv.conn,
			buffer:     make([]byte, len(data)),
			addr:       addr,
			readSignal: make(chan struct{}, 1),
		}
		copy(udpConnection.buffer, data)
		srv.udpConns[addr.String()] = udpConnection

		var err error
		session, err := kcp.NewConn3(0, addr, nil, 10, 3, udpConnection)
		if err != nil {
			log.Printf("Error creating KCP connection: %v", err)
			return
		}
		// session.SetStreamMode(true)
		session.SetNoDelay(1, 20, 2, 1)
		session.SetWindowSize(256, 256)
		session.SetMtu(1197)
		srv.kcpSessions[addr.String()] = session
		go srv.handleKCPSession(session)
	} else {
		udpConnection.mutex.Lock()
		udpConnection.buffer = make([]byte, len(data))
		copy(udpConnection.buffer, data)
		udpConnection.addr = addr
		udpConnection.mutex.Unlock()
	}
	// Signal that data is available
	udpConnection.readSignal <- struct{}{}
}

func (srv *server) handleUnreliablePacket(data []byte, addr *net.UDPAddr) {
	log.Printf("Received Unreliable packet from %s: %x\n", addr, data)
	// Further processing for raw UDP packets here
}

func (srv *server) handleKCPSession(session *kcp.UDPSession) {
	defer session.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := session.Read(buffer)
		if err != nil {
			log.Printf("Error reading from KCP connection: %v", err)
			return
		}

		fmt.Printf("Received %d bytes (KCP): %x\n", n, buffer[:n])
	}
}

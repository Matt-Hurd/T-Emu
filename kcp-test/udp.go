package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type UDPConn struct {
	Conn       *net.UDPConn
	Buffer     []byte
	Addr       *net.UDPAddr
	Mutex      sync.Mutex
	ReadSignal chan struct{}
	SendCount  int
	Start      time.Time
}

func (c *UDPConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	tempBuffer := make([]byte, 1500) // Adjust size as needed
	n, addr, err = c.Conn.ReadFromUDP(tempBuffer)
	if err != nil {
		return 0, nil, err
	}

	// Check if we have at least 3 bytes to strip off
	if n < 3 {
		return 0, nil, fmt.Errorf("received data is too short")
	}

	// Strip off the first 3 bytes and copy the rest into p
	copy(p, tempBuffer[:n])
	c.Buffer = tempBuffer[3:n]

	fmt.Printf("ReadFrom: %x\n", tempBuffer)

	return n, addr, nil
}

func (c *UDPConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	buf := make([]byte, 3)
	c.SendCount += 1
	buf[0] = 1
	buf[1] = byte(c.SendCount & 0xFF)
	buf[2] = byte((c.SendCount & 0xFF00) >> 8)
	buf = append(buf, p[8:]...)
	// log.Printf("Sent message: %x", buf)
	return c.Conn.WriteTo(buf, addr)
}

func (c *UDPConn) WriteToUnreliable(p []byte, addr net.Addr) (n int, err error) {
	buf := make([]byte, 3)
	c.SendCount += 1
	buf[0] = 2
	buf[1] = byte(c.SendCount & 0xFF)
	buf[2] = byte((c.SendCount & 0xFF00) >> 8)
	buf = append(buf, p...)
	return c.Conn.WriteTo(buf, addr)
}

func (c *UDPConn) Close() error {
	return c.Conn.Close()
}

func (c *UDPConn) LocalAddr() net.Addr {
	return c.Conn.LocalAddr()
}

func (c *UDPConn) SetDeadline(t time.Time) error {
	return c.Conn.SetDeadline(t)
}

func (c *UDPConn) SetReadDeadline(t time.Time) error {
	return c.Conn.SetReadDeadline(t)
}

func (c *UDPConn) SetWriteDeadline(t time.Time) error {
	return c.Conn.SetWriteDeadline(t)
}

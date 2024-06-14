package udp

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
}

func (c *UDPConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) {
	<-c.ReadSignal
	c.Mutex.Lock()
	defer c.Mutex.Unlock()

	if len(c.Buffer) > 0 {
		n = copy(p, c.Buffer)
		addr = c.Addr
		c.Buffer = nil
		fmt.Printf("Read %d bytes from %s: %x\n", n, addr, p[:n])
		return n, addr, nil
	}
	return 0, nil, nil
}

func (c *UDPConn) WriteTo(p []byte, addr net.Addr) (n int, err error) {
	buf := make([]byte, 3)
	c.SendCount += 1
	buf[0] = 1
	buf[1] = byte(c.SendCount & 0xFF)
	buf[2] = byte((c.SendCount & 0xFF00) >> 8)
	buf = append(buf, p[8:]...)
	fmt.Printf("Writing %d bytes to %s: %x\n", len(buf), addr, buf)
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

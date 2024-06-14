package main

import (
	"log"
	"net"
	"time"

	kcp "github.com/xtaci/kcp-go/v5"
)

func main() {

	udpaddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:9090")
	if err != nil {
		log.Fatalf("Error resolving UDP address: %v", err)
	}
	network := "udp4"
	if udpaddr.IP.To4() == nil {
		network = "udp"
	}

	conn, err := net.ListenUDP(network, nil)
	if err != nil {
		log.Fatalf("Error listening on UDP: %v", err)
	}

	// Connect to the KCP server
	kcpconn, err := kcp.NewConn3(0, udpaddr, nil, 20, 6, conn)
	if err != nil {
		log.Fatalf("Error connecting to KCP server: %v", err)
	}
	defer kcpconn.Close()

	//conn.SetStreamMode(true) // Set stream mode for testing
	kcpconn.SetNoDelay(1, 30, 3, 1)
	kcpconn.SetWindowSize(256, 256)
	kcpconn.SetMtu(1197)

	message := []byte("111")
	_, err = kcpconn.Write(message)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	log.Printf("Sent message: %s", message)

	// Wait to receive a response
	buffer := make([]byte, 1024)
	kcpconn.SetReadDeadline(time.Now().Add(5 * time.Second)) // Set a read deadline to avoid blocking indefinitely
	n, err := kcpconn.Read(buffer)
	if err != nil {
		log.Fatalf("Error reading from KCP connection: %v", err)
	}

	log.Printf("Received message: %x", buffer[:n])

	time.Sleep(1 * time.Second) // Give the server some time to print the message before exiting
}

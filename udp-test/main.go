package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("udp", "127.0.0.1:9090")
	if err != nil {
		log.Fatalf("Error connecting to UDP server: %v", err)
	}
	defer conn.Close()

	message := []byte("Hello, UDP!")
	_, err = conn.Write(message)
	if err != nil {
		log.Fatalf("Error sending message: %v", err)
	}

	log.Printf("Sent message: %s", message)
}

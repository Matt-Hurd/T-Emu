package main

import (
	"bytes"
	"encoding/hex"
	"game-server/models"
	"game-server/server"
	"strings"

	"github.com/rs/zerolog/log"
)

func main() {
	// test()
	srv := server.NewServer()
	err := srv.Start()
	if err != nil {
		log.Err(err)
	}
}

func parse_input(data []byte) {
	buffer := bytes.NewBuffer(data)
	for buffer.Len() > 0 {
		dataPacket := &models.DataPacket{}
		err := dataPacket.Parse(buffer)
		if err != nil {
			log.Err(err)
			break
		}
		// fmt.Printf("Parsed packet: %v\n", dataPacket)
	}
}

func hexStringToByteArray(s string) ([]byte, error) {
	parts := strings.Split(s, "-")

	bytes := make([]byte, len(parts))

	for i, part := range parts {
		byteVal, err := hex.DecodeString(part)
		if err != nil {
			return nil, err
		}
		bytes[i] = byteVal[0]
	}

	return bytes, nil
}

func test() {
	data := "07-00-02-00-FB-2A-2E-06-BB-01-0A"
	b, _ := hexStringToByteArray(data)
	parse_input(b)
}

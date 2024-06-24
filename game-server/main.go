package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
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
			panic(err)
		}
		fmt.Printf("data: %+v\n", dataPacket.GamePacket)
		newBuf := dataPacket.Write()
		// err = dataPacket.Parse(bytes.NewBuffer(newBuf))
		// if err != nil {
		// 	panic(err)
		// }
		fmt.Printf("written first 10 bytes of buffer: %x\n", newBuf[:10])
		if !bytes.Equal(newBuf, data) {
			fmt.Printf("Equals: %t\n", bytes.Equal(newBuf, data))
			diff, err := byteDiff(newBuf, data)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Last 10 bytes of original: %x\n", data[len(data)-10:])
			fmt.Printf("Last 10 bytes of new: %x\n", newBuf[len(newBuf)-10:])
			for i := range diff {
				if diff[i] != 0 && i < 100 {
					fmt.Printf("differs at %d, expected %d, got %d\n", i, data[i], newBuf[i])
				}
			}
		}
	}
}

func byteDiff(bs1, bs2 []byte) ([]int16, error) {
	if (bs1 == nil) || (bs2 == nil) {
		return nil, fmt.Errorf("expected a byte slice but got nil")
	}
	if len(bs1) != len(bs2) {
		fmt.Printf("mismatched lengths, %d != %d", len(bs1), len(bs2))
	}

	diff := make([]int16, len(bs2))
	for i := range bs2 {
		diff[i] = int16(bs1[i]) - int16(bs2[i])
	}
	return diff, nil
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
	data := ""
	b, _ := hexStringToByteArray(data)
	parse_input(b)

	// x := &response.WorldSpawn{}
	// x.GetDefault()
	// fmt.Printf("WorldSpawn: %v\n", x)
}

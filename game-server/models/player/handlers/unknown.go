package handlers

import (
	"fmt"
	"game-server/models"
	"game-server/models/game/enums/network"
)

func HandleUnknownPacket(packet *models.DataPacket, packet_1 *models.NetworkMessage) {
	fmt.Printf("Unhandled Packet Type: %s, Reliable(%t), Data: %x\n", network.GetPacketType(packet.GamePacketType), packet_1.Channel == models.NetworkChannelReliable, packet_1.Buffer[4:])
}

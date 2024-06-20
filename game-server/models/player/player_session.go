package player

import (
	"bytes"
	"fmt"
	"game-server/models"
	networkModels "game-server/models/game/enums/network"
	"game-server/models/game/request"
	"game-server/models/player/handlers"
	"game-server/network"
)

type PlayerSession struct {
	Connection *network.NetworkManager
	ProfileId  string
	Token      string
}

func (ps *PlayerSession) HandleDataPacket(packet *models.NetworkMessage) {
	dataPacket := &models.DataPacket{}
	buf := bytes.NewBuffer(packet.Buffer)
	for buf.Len() > 0 {
		err := dataPacket.Parse(buf)
		if err != nil {
			fmt.Println("Error parsing data packet:", err)
			return
		}
		switch dataPacket.GamePacketType {
		case uint16(networkModels.Command):
			handlers.HandlePacketCmdRequest(dataPacket.GamePacket.(*request.PacketCmdRequest), ps.Connection)
		case uint16(networkModels.Ready):
			handlers.HandlePacketPlayerReady(dataPacket.GamePacket.(*request.PacketClientReady), ps.Connection)
		case uint16(networkModels.ConnectionRequest):
			handlers.HandlePacketConnection(dataPacket.GamePacket.(*request.PacketConnection), ps.Connection, &ps.ProfileId, &ps.Token)
		case uint16(networkModels.ProgressReport):
			handlers.HandlePacketProgressReport(dataPacket.GamePacket.(*request.PacketProgressReport), ps.Connection)
		case uint16(networkModels.HLAPI):
			handlers.HandlePacketHLAPIRequest(dataPacket.GamePacket.(*request.PacketHLAPIRequest), ps.Connection, ps.ProfileId)
		default:
			handlers.HandleUnknownPacket(dataPacket, packet)
		}
	}
}

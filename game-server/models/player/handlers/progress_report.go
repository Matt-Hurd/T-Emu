package handlers

import (
	"fmt"
	"game-server/helpers"
	networkModels "game-server/models/game/enums/network"
	"game-server/models/game/request"
	"game-server/models/game/response"
	"game-server/network"
)

func HandlePacketProgressReport(packet *request.PacketProgressReport, g *network.NetworkManager) {
	fmt.Printf("Received ProgressReport: %v\n", packet)
	if packet.Progress == 1 {
		switch packet.Id {
		case 0:
			var err error
			nightmare := &response.NightMare{
				Id: 1,
			}
			nightmare.PrefabsData, err = helpers.CompressZlib([]byte(`[{"path":"assets/content/items/mods/barrels/barrel_mr43e-1c_510mm.bundle","rcid":"","FileName":"barrel_mr43e-1c_510mm"}]`))
			if err != nil {
				fmt.Println("Error compressing prefabs data:", err)
			}
			nightmare.CustomizationData, err = helpers.CompressZlib([]byte(`[]`))
			if err != nil {
				fmt.Println("Error compressing customization data:", err)
			}
			g.SendReliableDataPacket(networkModels.NightMare, nightmare)
		}
	}
}

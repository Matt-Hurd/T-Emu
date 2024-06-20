package handlers

import (
	"fmt"
	"game-server/helpers"
	networkModels "game-server/models/game/enums/network"
	"game-server/models/game/request"
	"game-server/models/game/response"
	"game-server/network"
	"time"
)

func HandlePacketConnection(packet *request.PacketConnection, g *network.NetworkManager, profileID *string, token *string) {
	var err error

	*profileID = packet.ProfileID
	*token = packet.Token

	connResponse := &response.PacketConnection{}
	connResponse.GetDefault()

	g.SendReliableDataPacket(networkModels.ConnectionRequest, connResponse)

	time.Sleep(10 * time.Second)

	nightMare := &response.NightMare{
		Id: 0,
	}
	nightMare.PrefabsData, err = helpers.CompressZlib([]byte(`[{"path":"assets/content/items/mods/barrels/barrel_pl15_izhmash_112mm_threaded_9x19.bundle","rcid":"","FileName":"barrel_pl15_izhmash_112mm_threaded_9x19"},{"path":"assets/content/items/mods/silencers/silencer_all_izhmash_pl15_std_9x19.bundle","rcid":"","FileName":"silencer_all_izhmash_pl15_std_9x19"}]`))
	if err != nil {
		fmt.Println("Error compressing prefabs data:", err)
	}
	nightMare.CustomizationData, err = helpers.CompressZlib([]byte(`["66043cc27502eca33a08cad0","5e9dc97c86f774054c19ac9a"]`))
	if err != nil {
		fmt.Println("Error compressing customization data:", err)
	}
	g.SendReliableDataPacket(networkModels.NightMare, nightMare)

}

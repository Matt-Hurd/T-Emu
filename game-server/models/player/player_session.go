package player

import "game-server/network"

type PlayerSession struct {
	Connection *network.NetworkManager
	ProfileId  string
	Token      string
}

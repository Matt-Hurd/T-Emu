package response

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
	playercommands "game-server/models/game/player_commands"
)

// 0xAC

type PacketSpawnObservedPlayers struct {
	Players map[byte]core.ObservedPlayer
}

func (p *PacketSpawnObservedPlayers) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteByte(buffer, byte(len(p.Players))); err != nil {
		return err
	}
	for key, value := range p.Players {
		if err := helpers.WriteByte(buffer, key); err != nil {
			return err
		}
		tmpBuffer := new(bytes.Buffer)
		if err := value.Serialize(tmpBuffer); err != nil {
			return err
		}
		if err := helpers.WriteBytesAndSize32(buffer, tmpBuffer.Bytes()); err != nil {
			return err
		}
	}
	return nil
}

func (p *PacketSpawnObservedPlayers) Deserialize(buffer *bytes.Buffer) error {
	var err error
	var count byte
	if err = helpers.ReadByte(buffer, &count); err != nil {
		return err
	}
	p.Players = make(map[byte]core.ObservedPlayer)
	for i := 0; i < int(count); i++ {
		var key byte
		if err = helpers.ReadByte(buffer, &key); err != nil {
			return err
		}
		var tmpBuffer []byte
		err := helpers.ReadBytesAndSize32(buffer, &tmpBuffer)
		if err != nil {
			return err
		}
		value := core.ObservedPlayer{
			HandsController: &playercommands.PlayerCommandSetHands{},
		}
		if err := value.Deserialize(bytes.NewBuffer(tmpBuffer)); err != nil {
			panic(err)
		}
		p.Players[key] = value
	}
	return nil
}

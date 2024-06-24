package response

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
	playercommands "game-server/models/game/player_commands"
)

// 0xAC

type PacketSpawnObservedPlayer struct {
	Id     byte
	Player core.ObservedPlayer
}

func (p *PacketSpawnObservedPlayer) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteByte(buffer, p.Id); err != nil {
		return err
	}
	tmpBuffer := new(bytes.Buffer)
	if err := p.Player.Serialize(tmpBuffer); err != nil {
		return err
	}
	if err := helpers.WriteBytesAndSize32(buffer, tmpBuffer.Bytes()); err != nil {
		return err
	}
	return nil
}

func (p *PacketSpawnObservedPlayer) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadByte(buffer, &p.Id); err != nil {
		return err
	}
	var tmpBuffer []byte
	err = helpers.ReadBytesAndSize32(buffer, &tmpBuffer)
	if err != nil {
		return err
	}
	p.Player = core.ObservedPlayer{
		HandsController: &playercommands.PlayerCommandSetHands{},
	}
	if err := p.Player.Deserialize(bytes.NewBuffer(tmpBuffer)); err != nil {
		panic(err)
	}
	return nil
}

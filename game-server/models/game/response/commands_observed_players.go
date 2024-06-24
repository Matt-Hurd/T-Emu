package response

//0xAE

import (
	"bytes"
	"game-server/helpers"
	playercommands "game-server/models/game/player_commands"
)

// GStruct286
type PacketCommandsObservedPlayers struct {
	WorldTime              float32
	ObservedPlayerCommands map[byte][]playercommands.PlayerCommand
}

func (p *PacketCommandsObservedPlayers) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.WriteFloat32(buffer, p.WorldTime)
	if err != nil {
		return err
	}
	err = helpers.WriteByte(buffer, byte(len(p.ObservedPlayerCommands)))
	if err != nil {
		return err
	}
	previousKey := byte(0)
	for key, value := range p.ObservedPlayerCommands {
		key -= previousKey
		err = helpers.WriteByte(buffer, key)
		if err != nil {
			return err
		}
		err = helpers.WriteByte(buffer, byte(len(value)))
		if err != nil {
			return err
		}
		for _, v := range value {
			err = v.Serialize(buffer)
			if err != nil {
				return err
			}
		}
		previousKey = key
	}
	return nil
}

func (p *PacketCommandsObservedPlayers) Deserialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.ReadFloat32(buffer, &p.WorldTime)
	if err != nil {
		return err
	}
	var count byte
	err = helpers.ReadByte(buffer, &count)
	if err != nil {
		return err
	}
	p.ObservedPlayerCommands = make(map[byte][]playercommands.PlayerCommand)
	previousKey := byte(0)
	for i := 0; i < int(count); i++ {
		var key byte
		err = helpers.ReadByte(buffer, &key)
		if err != nil {
			return err
		}
		key += previousKey
		var count2 byte
		err = helpers.ReadByte(buffer, &count2)
		if err != nil {
			return err
		}
		p.ObservedPlayerCommands[key] = make([]playercommands.PlayerCommand, count2)
		for j := 0; j < int(count2); j++ {
			err = p.ObservedPlayerCommands[key][j].Deserialize(buffer)
			if err != nil {
				return err
			}
		}
		previousKey = key
	}
	return nil
}

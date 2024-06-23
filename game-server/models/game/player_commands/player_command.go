package playercommands

import (
	"bytes"
	"fmt"
	"game-server/models/game/enums"
)

type PlayerCommand struct {
	Type enums.CommandMessageType
	Data PlayerCommandMsg
}

type PlayerCommandMsg interface {
	Serialize(buffer *bytes.Buffer) error
	Deserialize(buffer *bytes.Buffer) error
}

func (command *PlayerCommand) Serialize(buffer *bytes.Buffer) error {
	buffer.WriteByte(byte(command.Type))
	err := command.Data.Serialize(buffer)
	if err != nil {
		return err
	}
	return nil
}

func (command *PlayerCommand) Deserialize(buffer *bytes.Buffer) error {
	val, err := buffer.ReadByte()
	command.Type = enums.CommandMessageType(val)
	if err != nil {
		return err
	}
	switch command.Type {
	case 0:
		command.Data = &PlayerCommandDeath{}
	case 3:
		command.Data = &PlayerCommandSkillParameters{}
	case 4:
		command.Data = &PlayerCommandTemperature{}
	case 5:
		command.Data = &PlayerCommandHealthStatus{}
	case 6:
		command.Data = &PlayerCommandMedEffectStatus{}
	case 8:
		command.Data = &PlayerCommandPhysicalParameters{}
	default:
		return fmt.Errorf("unknown command type: %d", command.Type)
	}
	err = command.Data.Deserialize(buffer)
	if err != nil {
		return err
	}
	return nil
}

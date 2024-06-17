package core

import (
	"bytes"
	"game-server/helpers"
)

type PlayerStateLimits struct {
	MinSpeed float32
	MaxSpeed float32
}

func DeserializePlayerStateLimits(buffer *bytes.Buffer) (PlayerStateLimits, error) {
	var psl PlayerStateLimits
	var err error

	if err = helpers.ReadFloat32(buffer, &psl.MinSpeed); err != nil {
		return psl, err
	}
	if err = helpers.ReadFloat32(buffer, &psl.MaxSpeed); err != nil {
		return psl, err
	}
	return psl, nil
}

func (psl *PlayerStateLimits) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = helpers.WriteFloat32(buffer, psl.MinSpeed); err != nil {
		return err
	}
	if err = helpers.WriteFloat32(buffer, psl.MaxSpeed); err != nil {
		return err
	}
	return nil
}

func DeserializeToDict(buffer *bytes.Buffer) (map[EPlayerState]PlayerStateLimits, error) {
	dict := make(map[EPlayerState]PlayerStateLimits)
	var err error

	var count int32
	if err = helpers.ReadInt32(buffer, &count); err != nil {
		return nil, err
	}
	for i := 0; i < int(count); i++ {
		var key byte
		if err = helpers.ReadByte(buffer, &key); err != nil {
			return nil, err
		}
		var value PlayerStateLimits
		if value, err = DeserializePlayerStateLimits(buffer); err != nil {
			return nil, err
		}
		dict[EPlayerState(key)] = value
	}
	return dict, nil
}

func SerializeToDict(buffer *bytes.Buffer, dict map[EPlayerState]PlayerStateLimits) error {
	var err error

	count := int32(len(dict))
	if err = helpers.WriteInt32(buffer, count); err != nil {
		return err
	}
	for key, value := range dict {
		if err = helpers.WriteByte(buffer, byte(key)); err != nil {
			return err
		}
		if err = value.Serialize(buffer); err != nil {
			return err
		}
	}
	return nil
}

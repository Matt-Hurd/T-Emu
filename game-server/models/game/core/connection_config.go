package core

import (
	"bytes"
)

type EPlayerState int

type ConnectionConfig struct {
	DefaultPlayerStateLimits PlayerStateLimits
	DictPlayerStateLimits    map[EPlayerState]PlayerStateLimits
}

func (cc *ConnectionConfig) Deserialize(buffer *bytes.Buffer) error {
	var err error

	if cc.DefaultPlayerStateLimits, err = DeserializePlayerStateLimits(buffer); err != nil {
		return err
	}
	if cc.DictPlayerStateLimits, err = DeserializeToDict(buffer); err != nil {
		return err
	}
	return nil
}

func (cc *ConnectionConfig) Serialize(buffer *bytes.Buffer) error {
	var err error

	if err = cc.DefaultPlayerStateLimits.Serialize(buffer); err != nil {
		return err
	}
	if err = SerializeToDict(buffer, cc.DictPlayerStateLimits); err != nil {
		return err
	}
	return nil
}

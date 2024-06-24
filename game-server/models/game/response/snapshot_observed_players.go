package response

import (
	"bytes"
	"fmt"
	"game-server/helpers"
	"game-server/models/game/core"
)

// GStruct286
type PacketSnapshotObservedPlayers struct {
	WorldTime               float32
	ObservedPlayerSnapshots map[byte]*core.PlayerSnapshot
}

func (p *PacketSnapshotObservedPlayers) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.WriteFloat32(buffer, p.WorldTime)
	if err != nil {
		return err
	}
	err = helpers.WriteByte(buffer, byte(len(p.ObservedPlayerSnapshots)))
	if err != nil {
		return err
	}
	previousKey := byte(0)
	for key, value := range p.ObservedPlayerSnapshots {
		key -= previousKey
		err = helpers.WriteByte(buffer, key)
		if err != nil {
			return err
		}
		err = value.Serialize(buffer)
		if err != nil {
			return err
		}
		previousKey = key
	}
	return nil
}

func (p *PacketSnapshotObservedPlayers) Deserialize(buffer *bytes.Buffer) error {
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
	p.ObservedPlayerSnapshots = make(map[byte]*core.PlayerSnapshot)
	previousKey := byte(0)
	for i := 0; i < int(count); i++ {
		var key byte
		err = helpers.ReadByte(buffer, &key)
		if err != nil {
			return err
		}
		key += previousKey
		p.ObservedPlayerSnapshots[key] = &core.PlayerSnapshot{}
		err = p.ObservedPlayerSnapshots[key].Deserialize(buffer)
		if err != nil {
			return err
		}
		fmt.Printf("PlayerSnapshot: %v\n", p.ObservedPlayerSnapshots[key])
		previousKey = key
	}
	return nil
}

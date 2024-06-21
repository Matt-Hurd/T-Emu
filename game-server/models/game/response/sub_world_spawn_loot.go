package response

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
)

type SubWorldSpawnLoot struct {
	Flag  bool
	Loot  []byte
	Count int32
}

func (p *SubWorldSpawnLoot) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.WriteBool(buffer, p.Flag)
	if err != nil {
		return err
	}
	if p.Flag {
		err = helpers.WriteBytesAndSize(buffer, p.Loot)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *SubWorldSpawnLoot) Deserialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.ReadBool(buffer, &p.Flag)
	if err != nil {
		return err
	}
	if p.Flag {
		err = helpers.ReadBytesAndSize(buffer, &p.Loot)
		if err != nil {
			return err
		}
		if unzipped, err := helpers.DecompressZlib(p.Loot); err != nil {
			return err
		} else {
			newBuf := bytes.NewBuffer(unzipped)
			err = helpers.ReadInt32(newBuf, &p.Count)
			if err != nil {
				return err
			}
			for i := 0; i < int(p.Count); i++ {
				var loot core.Serializable
				err = core.ReadPolymorph(newBuf, &loot)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

package response

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
)

type SubWorldSpawnLoot struct {
	Flag      bool
	LootItems []core.Serializable //core.LootItem
}

func (p *SubWorldSpawnLoot) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.WriteBool(buffer, p.Flag)
	if err != nil {
		return err
	}
	if p.Flag {
		tmpBuf := new(bytes.Buffer)
		err = helpers.WriteInt32(tmpBuf, int32(len(p.LootItems)))
		if err != nil {
			return err
		}
		for _, v := range p.LootItems {
			err = core.WritePolymorph(tmpBuf, v)
			if err != nil {
				return err
			}
		}
		zipped, err := helpers.CompressZlib(tmpBuf.Bytes())
		if err != nil {
			return err
		}
		err = helpers.WriteBytesAndSize(buffer, zipped)
		if err != nil {
			return err
		}
	} else {
		err = helpers.WriteBytesAndSize(buffer, []byte{})
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
		var tmp []byte
		err = helpers.ReadBytesAndSize(buffer, &tmp)
		if err != nil {
			return err
		}
		if unzipped, err := helpers.DecompressZlib(tmp); err != nil {
			return err
		} else {
			newBuf := bytes.NewBuffer(unzipped)
			count := int32(0)
			err = helpers.ReadInt32(newBuf, &count)
			if err != nil {
				return err
			}
			p.LootItems = make([]core.Serializable, count)
			for i := 0; i < int(count); i++ {
				var loot core.Serializable
				err = core.ReadPolymorph(newBuf, &loot)
				if err != nil {
					return err
				}
				p.LootItems[i] = loot
			}
		}
	}
	return nil
}

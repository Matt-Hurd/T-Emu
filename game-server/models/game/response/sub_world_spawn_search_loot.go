package response

import (
	"bytes"
	"game-server/helpers"
	"game-server/models/game/core"
)

type SubWorldSpawnSearchLoot struct {
	Flag       bool
	SearchLoot []core.SearchLoot
}

func (p *SubWorldSpawnSearchLoot) Serialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.WriteBool(buffer, p.Flag)
	if err != nil {
		return err
	}
	tmpBuf := new(bytes.Buffer)
	packedWriter := helpers.NewLimitedWriter(tmpBuf)
	err = packedWriter.Write(int32(len(p.SearchLoot)))
	if err != nil {
		return err
	}
	for _, v := range p.SearchLoot {
		err = v.Serialize(packedWriter)
		if err != nil {
			return err
		}
	}
	packedWriter.FlushBits()
	err = helpers.WriteBytesAndSize(buffer, tmpBuf.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (p *SubWorldSpawnSearchLoot) Deserialize(buffer *bytes.Buffer) error {
	var err error
	err = helpers.ReadBool(buffer, &p.Flag)
	if err != nil {
		return err
	}
	var tmp []byte
	err = helpers.ReadBytesAndSize(buffer, &tmp)
	if err != nil {
		return err
	}
	tmpBuf := bytes.NewBuffer(tmp)
	packedReader := helpers.NewLimitedReader(tmpBuf)
	count := int32(0)
	err = packedReader.Read(&count)
	if err != nil {
		return err
	}
	p.SearchLoot = make([]core.SearchLoot, count)
	for i := 0; i < int(count); i++ {
		err = p.SearchLoot[i].Deserialize(packedReader)
		if err != nil {
			return err
		}
	}
	return nil
}

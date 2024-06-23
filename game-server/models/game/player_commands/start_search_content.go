package playercommands

import (
	"bytes"
	"game-server/helpers"
)

type PlayerCommandStartSearchContent struct {
	SearchSoundName string
}

func (msg *PlayerCommandStartSearchContent) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteString(buffer, msg.SearchSoundName); err != nil {
		return err
	}
	return nil
}

func (msg *PlayerCommandStartSearchContent) Deserialize(buffer *bytes.Buffer) error {
	var err error
	if err = helpers.ReadString(buffer, &msg.SearchSoundName); err != nil {
		return err
	}
	return nil
}

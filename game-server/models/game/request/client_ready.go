package request

import (
	"bytes"
)

type PacketClientReady struct {
}

func (p *PacketClientReady) Deserialize(buffer *bytes.Buffer) error {
	return nil
}

func (p *PacketClientReady) Serialize(buffer *bytes.Buffer) error {
	return nil
}

package response

import (
	"bytes"
)

type PacketHLAPIResponse struct {
	Data []byte
}

func (p *PacketHLAPIResponse) Deserialize(buffer *bytes.Buffer) error {
	p.Data = make([]byte, buffer.Len())
	if _, err := buffer.Read(p.Data); err != nil {
		return err
	}
	return nil
}

func (p *PacketHLAPIResponse) Serialize(buffer *bytes.Buffer) error {
	if _, err := buffer.Write(p.Data); err != nil {
		return err
	}
	return nil
}

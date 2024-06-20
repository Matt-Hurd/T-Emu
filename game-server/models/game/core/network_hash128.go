package core

import "bytes"

type NetworkHash128 struct {
	i0  byte
	i1  byte
	i2  byte
	i3  byte
	i4  byte
	i5  byte
	i6  byte
	i7  byte
	i8  byte
	i9  byte
	i10 byte
	i11 byte
	i12 byte
	i13 byte
	i14 byte
	i15 byte
}

func (n *NetworkHash128) Deserialize(buffer *bytes.Buffer) error {
	tmp := make([]byte, 16)
	_, err := buffer.Read(tmp)
	if err != nil {
		return err
	}
	n.i0 = tmp[0]
	n.i1 = tmp[1]
	n.i2 = tmp[2]
	n.i3 = tmp[3]
	n.i4 = tmp[4]
	n.i5 = tmp[5]
	n.i6 = tmp[6]
	n.i7 = tmp[7]
	n.i8 = tmp[8]
	n.i9 = tmp[9]
	n.i10 = tmp[10]
	n.i11 = tmp[11]
	n.i12 = tmp[12]
	n.i13 = tmp[13]
	n.i14 = tmp[14]
	n.i15 = tmp[15]
	return nil
}

func (n *NetworkHash128) Serialize(buffer *bytes.Buffer) error {
	buffer.WriteByte(n.i0)
	buffer.WriteByte(n.i1)
	buffer.WriteByte(n.i2)
	buffer.WriteByte(n.i3)
	buffer.WriteByte(n.i4)
	buffer.WriteByte(n.i5)
	buffer.WriteByte(n.i6)
	buffer.WriteByte(n.i7)
	buffer.WriteByte(n.i8)
	buffer.WriteByte(n.i9)
	buffer.WriteByte(n.i10)
	buffer.WriteByte(n.i11)
	buffer.WriteByte(n.i12)
	buffer.WriteByte(n.i13)
	buffer.WriteByte(n.i14)
	buffer.WriteByte(n.i15)
	return nil
}

func (n *NetworkHash128) FromBytes(data []byte) {
	n.i0 = data[0]
	n.i1 = data[1]
	n.i2 = data[2]
	n.i3 = data[3]
	n.i4 = data[4]
	n.i5 = data[5]
	n.i6 = data[6]
	n.i7 = data[7]
	n.i8 = data[8]
	n.i9 = data[9]
	n.i10 = data[10]
	n.i11 = data[11]
	n.i12 = data[12]
	n.i13 = data[13]
	n.i14 = data[14]
	n.i15 = data[15]
}

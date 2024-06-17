package response

import (
	"bytes"
	"fmt"

	"game-server/helpers"
)

type NightMare struct {
	Id                int32
	PrefabsData       []byte
	CustomizationData []byte
}

func (nm *NightMare) Deserialize(buffer *bytes.Buffer) error {
	if err := helpers.ReadInt32(buffer, &nm.Id); err != nil {
		return err
	}

	if err := helpers.ReadBytesAndSize(buffer, &nm.PrefabsData); err != nil {
		return err
	}
	if err := helpers.ReadBytesAndSize(buffer, &nm.CustomizationData); err != nil {
		return err
	}

	return nil
}

func (nm *NightMare) Serialize(buffer *bytes.Buffer) error {
	if err := helpers.WriteInt32(buffer, nm.Id); err != nil {
		return err
	}

	if err := helpers.WriteBytesAndSize(buffer, nm.PrefabsData); err != nil {
		return err
	}

	if err := helpers.WriteBytesAndSize(buffer, nm.CustomizationData); err != nil {
		return err
	}

	return nil
}

func (nm *NightMare) DebugPrintAll() {
	fmt.Printf("Id: %d\n", nm.Id)
	d, err := Decompress(nm.PrefabsData)
	if err != nil {
		fmt.Printf("Decompress PrefabsData error: %v\n", err)
	}
	fmt.Printf("PrefabsData: %s\n", d)
	d, err = Decompress(nm.CustomizationData)
	if err != nil {
		fmt.Printf("Decompress CustomizationData error: %v\n", err)
	}
	fmt.Printf("CustomizationData: %s\n", d)
}

func Decompress(data []byte) ([]byte, error) {
	return helpers.DecompressZlib(data)
}

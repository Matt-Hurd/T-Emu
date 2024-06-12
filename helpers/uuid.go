package helpers

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateUUID() string {
	bytes := make([]byte, 12)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

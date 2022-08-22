package services

import (
	"crypto/sha256"
	"fmt"
)

func SHA256ENCODER(text string) string {
	str_encoder := sha256.Sum256([]byte(text))

	return fmt.Sprintf("%x", str_encoder)
}

package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func HashPassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))
	hash := hex.EncodeToString(h.Sum(nil))
	return fmt.Sprint(hash)
}

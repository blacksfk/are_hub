package hash

import (
	"crypto/rand"
	"encoding/base64"
)

const (
	KEY_LEN = 16
)

func GenerateKey() (string, error) {
	// create and fill the byte slice with secure, random bytes
	key := make([]byte, KEY_LEN)
	_, e := rand.Read(key)

	if e != nil {
		return "", e
	}

	return base64.StdEncoding.EncodeToString(key), nil
}

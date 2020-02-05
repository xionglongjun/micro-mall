package utils

import (
	"encoding/base64"
	"log"

	"golang.org/x/crypto/scrypt"
)

// EncodeSalt ...
func EncodeSalt(password, salt string) (string, error) {
	dk, err := scrypt.Key([]byte(password), []byte(salt), 1<<15, 8, 1, 32)
	if err != nil {
		log.Fatal(err)
		return password, err
	}
	return base64.StdEncoding.EncodeToString(dk), nil
}

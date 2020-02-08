package utils

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"time"

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

// GenValidateCode ...
func GenValidateCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

// ValidateMobile ...
func ValidateMobile(mobile string) bool {
	ok, _ := regexp.MatchString(`^(1[1-9]\d{9})$`, mobile)
	return ok
}

// EncodeMD5 md5 encryption
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

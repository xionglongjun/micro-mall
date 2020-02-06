package utils

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func ValidateMobile(mobile string) bool {
	ok, _ := regexp.MatchString(`^(1[1-9]\d{9})$`, mobile)
	return ok
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

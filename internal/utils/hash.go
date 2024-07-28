package utils

import (
	"crypto/sha256"
	"fmt"
)

var (
	algo = sha256.New()
)

func HashString(s string) string {
	algo.Reset()
	algo.Write([]byte(s))
	return fmt.Sprintf("%x", algo.Sum(nil))
}

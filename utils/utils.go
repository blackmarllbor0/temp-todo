package utils

import (
	"crypto/rand"
	"fmt"
	"log"
)

// GenerateId генерирует новый id
func GenerateId() string {
	b := make([]byte, 16)

	if _, err := rand.Read(b); err != nil {
		log.Fatal(err.Error())
	}

	return fmt.Sprintf("%x", b)
}

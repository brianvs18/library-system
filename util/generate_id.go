package util

import (
	"crypto/rand"
	"log"
)

func GenerateId() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, 10)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal(err)
	}

	for i := range bytes {
		bytes[i] = charset[bytes[i]%byte(len(charset))]
	}

	return string(bytes)
}

package util

import (
	"math/rand"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digitBytes  = "0123456789"
	totalBytes  = letterBytes + digitBytes
)

func GenerateTemporaryPassword(length int) string {
	if length <= 0 {
		return ""
	}

	rand.Seed(time.Now().UnixNano())

	password := make([]byte, length)

	password[0] = digitBytes[rand.Intn(len(digitBytes))]

	for i := 1; i < length; i++ {
		password[i] = totalBytes[rand.Intn(len(totalBytes))]
	}

	rand.Shuffle(length, func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password)
}

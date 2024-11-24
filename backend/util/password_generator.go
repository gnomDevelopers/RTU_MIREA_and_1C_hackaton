package util

import (
	"math/rand"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

// GenerateTemporaryPassword генерирует временный пароль заданной длины
func GenerateTemporaryPassword(length int) string {
	rand.Seed(time.Now().UnixNano())

	// Создаем срез для пароля
	password := make([]byte, length)
	for i := 0; i < length; i++ {
		password[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(password)
}

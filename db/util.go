package db

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const SaltSize = 16
const StringSize = 32

func HashPassword(pw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.DefaultCost)
	return string(hash), err
}

func DoesPasswordMatch(hashedPw, pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPw), []byte(pw))
	if err == nil {
		return true
	}
	return false
}

func RandomString() string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, StringSize)
	for i := 0; i < StringSize; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}

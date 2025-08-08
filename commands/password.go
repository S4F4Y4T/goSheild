package commands

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+[]{}<>?"

func generateRandomPassword(length int) (string, error) {
	password := make([]byte, length)
	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[num.Int64()]
	}
	return string(password), nil
}

func passwordGenerate() {
	fmt.Println("Generating secure password...")
	password, err := generateRandomPassword(16)
	if err != nil {
		fmt.Println("Error generating password:", err)
		return
	}
	fmt.Println("Generated Password:", password)
}
package main

import (
	"fmt"

	"go-config-encryption/pkg/crypto/aesCbc256"
	"go-config-encryption/pkg/crypto/hash"
)

func main() {
	plainText := "plain text will be encrypted"
	secretKey := "3s6v9y$B&E)H@McQfTjWnZq4t7w!z%C*"

	encrypted, _ := aesCbc256.Encrypt(plainText, secretKey, "")
	decrypted, _ := aesCbc256.Decrypt(encrypted, secretKey, "")
	sha256 := hash.HashingSHA256(secretKey)

	fmt.Println("=== AES-256-CBC Encryption ===")
	fmt.Println("Encrypted Value: ", encrypted)
	fmt.Println("Decrypted Value: ", decrypted)

	fmt.Println("=== Hashing ===")
	fmt.Println("Hash SHA256 Value: ", sha256)
}

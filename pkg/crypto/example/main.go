package main

import (
	"fmt"

	"go-config-encryption/pkg/crypto/aesCbc256"
	"go-config-encryption/pkg/crypto/base64"
)

func main() {
	plainText := "plain text will be encrypted"
	secretKey := "3s6v9y$B&E)H@McQfTjWnZq4t7w!z%C*"

	encrypted, _ := aesCbc256.Encrypt(plainText, secretKey, "")
	decrypted, _ := aesCbc256.Decrypt(encrypted, secretKey, "")
	md5b64 := base64.GenerateMD5WithBase64Encoding(secretKey)

	fmt.Println("=== AES-256-CBC Encryption ===")
	fmt.Println("Encrypted Value: ", encrypted)
	fmt.Println("Decrypted Value: ", decrypted)

	fmt.Println("=== Base 64 Encoding ===")
	fmt.Println("MD5 Value: ", md5b64)
}

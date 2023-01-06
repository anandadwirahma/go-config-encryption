package base64

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

func GenerateMD5WithBase64Encoding(message string) string {
	b64Msg := GenerateBase64Encoding(message)

	hash := md5.New()
	hash.Write([]byte(b64Msg))

	return hex.EncodeToString(hash.Sum(nil))
}

func GenerateBase64Encoding(message string) string {
	return base64.StdEncoding.EncodeToString([]byte(message))
}

func DecodeBase64ToString(message string) string {
	s, _ := base64.StdEncoding.DecodeString(message)
	return string(s)
}

package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"

	"go-config-encryption/pkg/crypto/aesCbc256"
	"go-config-encryption/pkg/crypto/hash"
	"go-config-encryption/pkg/regex"
)

func New(cfg interface{}) (err error) {
	err = godotenv.Load()
	if err != nil {
		return
	}

	err = envconfig.Process("", cfg)
	if err != nil {
		return
	}

	return
}

type EncryptedValue string

func (s *EncryptedValue) Decode(cfgValue string) error {
	if !regex.RegexEncryptedValue.MatchString(cfgValue) {
		*s = EncryptedValue(cfgValue)
		return nil
	}

	var (
		secretKey = regex.ExtractEncryptedValue(os.Getenv("SECRETKEY"))
		encrypted = regex.ExtractEncryptedValue(cfgValue)
	)

	key := hash.HashingSHA256(secretKey)

	decrypted, err := aesCbc256.Decrypt(encrypted, key[32:], "")
	if err != nil {
		return err
	}

	*s = EncryptedValue(decrypted)

	return nil
}

func (s *EncryptedValue) String() string {
	return string(*s)
}

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"go-config-encryption/pkg/config"
)

type Config struct {
	Application Application `json:"application"`
	Database    Database    `json:"database"`
}

type Database struct {
	Name     config.EncryptedValue `json:"name"`
	Username config.EncryptedValue `json:"username"`
	Password config.EncryptedValue `json:"password"`
	Host     string                `json:"host"`
	Port     int                   `json:"port"`
}

type Application struct {
	ClientID  config.EncryptedValue `json:"client_id"`
	SecretKey config.EncryptedValue `json:"secret_key"`
}

func main() {
	cfg := &Config{}
	err := config.New(cfg)
	if err != nil {
		log.Fatalf("Failed read config file: %s", err.Error())
	}

	b, _ := json.MarshalIndent(cfg, "", "  ")

	fmt.Println(string(b))
}

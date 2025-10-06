package config

import (
	"flag"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	DatabaseType string
	BaseUrl      string
}

func NewConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("failed to load .env file")
	}
	baseUrl := os.Getenv("BASE_URL")
	if baseUrl == "" {
		baseUrl = "http://localhost:8080/"
	}

	storageType := flag.String("storage", "postgresql", "Тип хранилища (in-memory или postgresql)")
	flag.Parse()

	if *storageType != "postgresql" && *storageType != "in-memory" {
		panic("only postgresql and in-memory storages are supported")
	}
	return &Config{
		DatabaseType: *storageType,
		BaseUrl:      baseUrl,
	}

}

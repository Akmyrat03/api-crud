package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	App App
}

type App struct {
	Port string
}

var once sync.Once
var cfg Config

func LoadConfig() *Config {
	once.Do(func() {
		_ = godotenv.Load()

		cfg = Config{
			App: App{
				Port: getEnv("APP_PORT", ":8080"),
			},
		}
	})

	log.Printf("config = %v", cfg)

	return &cfg
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

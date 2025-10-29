package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	App      App
	Postgres PostgresConfig
}

type App struct {
	Port string
}

type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DB       string
	SslMode  string
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
			Postgres: PostgresConfig{
				User:     getEnv("POSTGRES_USER", ""),
				Password: getEnv("POSTGRES_PASSWORD", ""),
				Host:     getEnv("POSTGRES_HOST", ""),
				Port:     getEnv("POSTGRES_PORT", ""),
				DB:       getEnv("POSTGRES_DB", ""),
				SslMode:  getEnv("POSTGRES_SSLMODE", ""),
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

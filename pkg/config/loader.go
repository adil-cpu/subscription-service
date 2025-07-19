package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// храним параметры, прочитанные из .env или окружения
type Config struct {
	Port        string
	PostgresDSN string
	CorsOrigins string
}

// загржаем .env и собираем все переменные в Config
func Load() *Config {
	// читаем .env файл
	if err := godotenv.Load(); err != nil {
		log.Println(".env не найден, читаем переменные из окружения")
	}

	// по умолчанию ":8080"
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	return &Config{
		Port:        port,
		PostgresDSN: os.Getenv("POSTGRES_DSN"),
		CorsOrigins: os.Getenv("CORS_ORIGINS"),
	}
}

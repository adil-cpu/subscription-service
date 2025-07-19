// pkg/config/loader.go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config хранит параметры, прочитанные из .env или окружения
type Config struct {
	Port        string // Порт для HTTP-сервера, например ":8080"
	PostgresDSN string // Строка подключения к Postgres
	CorsOrigins string // Список разрешённых источников для CORS
}

// Load пытается загрузить .env и собирает все переменные в Config
func Load() *Config {
	// Пытаемся прочитать файл .env в корне проекта
	if err := godotenv.Load(); err != nil {
		log.Println("⚠ .env не найден, читаем переменные из окружения")
	}

	// Если PORT не задан, по умолчанию ":8080"
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	return &Config{
		Port:        port,
		PostgresDSN: os.Getenv("POSTGRES_DSN"), // ожидаем формат postgres://user:pass@host:port/db?sslmode=disable
		CorsOrigins: os.Getenv("CORS_ORIGINS"), // например "http://localhost:3000"
	}
}

package main

import (
	"github.com/adil-cpu/subscription-service/api"
	"github.com/adil-cpu/subscription-service/pkg/config"
	"github.com/adil-cpu/subscription-service/pkg/database"
	"github.com/adil-cpu/subscription-service/pkg/logger"
	"github.com/adil-cpu/subscription-service/pkg/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	// Swagger
	_ "github.com/adil-cpu/subscription-service/docs" // авто-сгенерированная документация
	swaggerFiles "github.com/swaggo/files"            // встроенные swagger файлы
	ginSwagger "github.com/swaggo/gin-swagger"        // хендлер для Swagger UI
)

// @title Subscription API
// @version 1.0
// @description API для управления подписками
// @host localhost:8080
// @BasePath /
//
// @contact.name Adil CPU
// @contact.url https://github.com/adil-cpu
// @contact.email example@example.com
func main() {
	cfg := config.Load()
	log := logger.New()

	db := database.Connect(cfg)

	// автоматически создаём таблицу, если её нет
	if err := db.AutoMigrate(&models.Subscription{}); err != nil {
		log.Fatal("ошибка миграции", zap.Error(err))
	}

	r := gin.New()

	// middleware: логгер и восстановление после паники
	r.Use(
		logger.GinMiddleware(log),
		logger.RecoveryMiddleware(log),
	)

	// роуты API
	api.RegisterRoutes(r, db)

	// Swagger UI по адресу http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Info("запуск сервера", zap.String("port", cfg.Port))
	if err := r.Run(cfg.Port); err != nil {
		log.Fatal("не удалось запустить сервер", zap.Error(err))
	}
}

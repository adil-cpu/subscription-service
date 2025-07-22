package main

import (
	"github.com/adil-cpu/subscription-service/api"
	"github.com/adil-cpu/subscription-service/pkg/config"
	"github.com/adil-cpu/subscription-service/pkg/database"
	"github.com/adil-cpu/subscription-service/pkg/logger"
	"github.com/adil-cpu/subscription-service/pkg/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()
	log := logger.New()

	db := database.Connect(cfg)

	if err := db.AutoMigrate(&models.Subscription{}); err != nil {
		log.Fatal("ошибка миграции", zap.Error(err))
	}

	r := gin.New()

	r.Use(
		logger.GinMiddleware(log),
		logger.RecoveryMiddleware(log),
	)

	// роуты
	api.RegisterRoutes(r)

	log.Info("запуск сервера", zap.String("port", cfg.Port))
	if err := r.Run(cfg.Port); err != nil {
		log.Fatal("не удалось запустить сервер", zap.Error(err))
	}
}

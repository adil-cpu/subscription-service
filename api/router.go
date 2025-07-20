package api

import (
	"github.com/adil-cpu/subscription-service/api/handler"
	"github.com/gin-gonic/gin"
)

// регистрируем все публичные маршруты API
func RegisterRoutes(r *gin.Engine) {
	r.GET("/healthz", handler.HealthHandler)
}

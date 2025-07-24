package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Проверка работоспособности
// @Description Возвращает 200 OK, если сервер работает
// @Tags health
// @Success 200 {object} map[string]string
// @Router /healthz [get]
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

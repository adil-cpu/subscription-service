package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// проверяем жив ли сервер
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

package api

import (
	"github.com/adil-cpu/subscription-service/api/handler"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	r.GET("/subscriptions", handler.GetAllSubscriptionsHandler(db))
	r.POST("/subscriptions", handler.CreateSubscriptionHandler(db))
	r.PUT("/subscriptions/:id", handler.UpdateSubscriptionHandler(db))
	r.DELETE("/subscriptions/:id", handler.DeleteSubscriptionHandler(db))

}

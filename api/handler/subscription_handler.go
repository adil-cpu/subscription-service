package handler

import (
	"net/http"
	"time"

	"github.com/adil-cpu/subscription-service/pkg/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// структура POST запроса
type CreateSubscriptionRequest struct {
	UserID    uint   `json:"user_id" binding:"required"`
	Plan      string `json:"plan" binding:"required"`
	StartDate string `json:"start_date"` // ISO 8601, пример: "2025-07-22T00:00:00Z"
	EndDate   string `json:"end_date"`
}

// POST /subscriptions — создаем новую подписку
func CreateSubscriptionHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req CreateSubscriptionRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат запроса"})
			return
		}

		start, err := time.Parse(time.RFC3339, req.StartDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат даты начала"})
			return
		}

		end, err := time.Parse(time.RFC3339, req.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат даты окончания"})
			return
		}

		sub := models.Subscription{
			UserID:    req.UserID,
			Plan:      req.Plan,
			StartDate: start,
			EndDate:   end,
		}

		if err := db.Create(&sub).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать подписку"})
			return
		}

		c.JSON(http.StatusCreated, sub)
	}
}

// GET /subscriptions — получаем все подписки
func GetAllSubscriptionsHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var subs []models.Subscription

		if err := db.Find(&subs).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить подписки"})
			return
		}

		c.JSON(http.StatusOK, subs)
	}
}

// PUT /subscriptions/:id — обновляем подписку по ID
func UpdateSubscriptionHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		var subscription models.Subscription
		if err := db.First(&subscription, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Подписка не найдена"})
			return
		}

		var input models.Subscription
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат данных"})
			return
		}

		subscription.Plan = input.Plan
		subscription.StartDate = input.StartDate
		subscription.EndDate = input.EndDate

		if err := db.Save(&subscription).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении подписки"})
			return
		}

		c.JSON(http.StatusOK, subscription)
	}
}

// DELETE /subscriptions/:id — удаляем подписку по ID
func DeleteSubscriptionHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if err := db.Delete(&models.Subscription{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при удалении подписки"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Подписка удалена"})
	}
}

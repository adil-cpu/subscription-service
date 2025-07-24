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
	UserID    uint    `json:"user_id" binding:"required"`
	Plan      string  `json:"plan" binding:"required"`
	Price     float64 `json:"price" binding:"required"`
	StartDate string  `json:"start_date"` // ISO 8601
	EndDate   string  `json:"end_date"`
}

// CreateSubscriptionHandler godoc
// @Summary Создать подписку
// @Description Создает новую подписку
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param subscription body CreateSubscriptionRequest true "Данные подписки"
// @Success 201 {object} models.Subscription
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions [post]
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
			Price:     req.Price,
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

// GetAllSubscriptionsHandler godoc
// @Summary Получить все подписки
// @Description Возвращает список всех подписок
// @Tags subscriptions
// @Produce json
// @Success 200 {array} models.Subscription
// @Failure 500 {object} map[string]string
// @Router /subscriptions [get]
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

// UpdateSubscriptionHandler godoc
// @Summary Обновить подписку
// @Description Обновляет данные подписки по ID
// @Tags subscriptions
// @Accept json
// @Produce json
// @Param id path int true "ID подписки"
// @Param subscription body models.Subscription true "Обновлённые данные"
// @Success 200 {object} models.Subscription
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions/{id} [put]
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

// DeleteSubscriptionHandler godoc
// @Summary Удалить подписку
// @Description Удаляет подписку по ID
// @Tags subscriptions
// @Produce json
// @Param id path int true "ID подписки"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /subscriptions/{id} [delete]
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

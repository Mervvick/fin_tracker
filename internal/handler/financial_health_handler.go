package handler

import (
	"fin_tracker/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FinancialHealthHandler struct {
	healthService *service.FinancialHealthService
}

func NewFinancialHealthHandler(hs *service.FinancialHealthService) *FinancialHealthHandler {
	return &FinancialHealthHandler{healthService: hs}
}

// GetFinancialHealth возвращает оценку финансового здоровья
// @Summary Финансовое здоровье пользователя
// @Tags analytics
// @Security BearerAuth
// @Produce json
// @Success 200 {object} model.FinancialHealth
// @Failure 500 {object} map[string]string
// @Router /analytics/financial-health [get]
func (h *FinancialHealthHandler) GetFinancialHealth(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	health, err := h.healthService.GetFinancialHealth(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to calculate financial health"})
		return
	}

	c.JSON(http.StatusOK, health)
}

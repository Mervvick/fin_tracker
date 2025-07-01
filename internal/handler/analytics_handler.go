package handler

import (
	"net/http"
	"time"

	"fin_tracker/internal/service"

	"github.com/gin-gonic/gin"
)

type AnalyticsHandler struct {
	analyticsService *service.AnalyticsService
}

func NewAnalyticsHandler(as *service.AnalyticsService) *AnalyticsHandler {
	return &AnalyticsHandler{analyticsService: as}
}

type SpendingByCategoryRequest struct {
	StartDate string `form:"start_date" binding:"required,datetime=2006-01-02"`
	EndDate   string `form:"end_date" binding:"required,datetime=2006-01-02"`
}

// GetSpendingByCategory возвращает распределение расходов по категориям
// @Summary Анализ расходов по категориям
// @Tags analytics
// @Security BearerAuth
// @Produce json
// @Param start_date query string true "Начальная дата (YYYY-MM-DD)"
// @Param end_date query string true "Конечная дата (YYYY-MM-DD)"
// @Success 200 {array} model.CategorySpending
// @Failure 400 {object} map[string]string
// @Router /analytics/spending-by-category [get]
func (h *AnalyticsHandler) GetSpendingByCategory(c *gin.Context) {
	var req SpendingByCategoryRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format"})
		return
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format"})
		return
	}

	results, err := h.analyticsService.GetSpendingByCategory(userID.(string), startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get analytics"})
		return
	}

	c.JSON(http.StatusOK, results)
}

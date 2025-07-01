package service

import (
	"fin_tracker/internal/model"
	"fin_tracker/internal/repository"
	"time"
)

type AnalyticsService struct {
	transactionRepo *repository.TransactionRepository
}

func NewAnalyticsService(tr *repository.TransactionRepository) *AnalyticsService {
	return &AnalyticsService{transactionRepo: tr}
}

func (s *AnalyticsService) GetSpendingByCategory(userID string, start, end time.Time) ([]model.CategorySpending, error) {
	results, err := s.transactionRepo.GetSpendingByCategory(userID, start, end)
	if err != nil {
		return nil, err
	}

	// Рассчитываем общую сумму расходов
	var totalSpent float64
	for _, item := range results {
		totalSpent += item.TotalAmount
	}

	// Вычисляем проценты для каждой категории
	for i := range results {
		if totalSpent != 0 {
			results[i].Percentage = (results[i].TotalAmount / totalSpent) * 100
		} else {
			results[i].Percentage = 0
		}
		// Делаем сумму положительной для отображения
		results[i].TotalAmount = -results[i].TotalAmount
	}

	return results, nil
}

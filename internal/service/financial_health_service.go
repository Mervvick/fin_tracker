package service

import (
	"fin_tracker/internal/model"
	"fin_tracker/internal/repository"
	"math"
)

type FinancialHealthService struct {
	transactionRepo *repository.TransactionRepository
}

func NewFinancialHealthService(tr *repository.TransactionRepository) *FinancialHealthService {
	return &FinancialHealthService{transactionRepo: tr}
}

func (s *FinancialHealthService) GetFinancialHealth(userID string) (*model.FinancialHealth, error) {
	totalBalance, income, expenses, largestCat, monthlyTrends, err := s.transactionRepo.GetFinancialHealthData(userID)
	if err != nil {
		return nil, err
	}

	// Рассчитываем соотношение доходов и расходов
	ratio := 0.0
	if expenses != 0 {
		ratio = math.Abs(income / expenses)
	}

	// Рассчитываем оценку финансового здоровья
	healthScore := calculateHealthScore(totalBalance, ratio)

	return &model.FinancialHealth{
		TotalBalance:       totalBalance,
		IncomeExpenseRatio: ratio,
		LargestExpenseCat:  largestCat,
		MonthlyTrends:      monthlyTrends,
		HealthScore:        healthScore,
	}, nil
}

func calculateHealthScore(balance, ratio float64) int {
	// Простая эвристика для оценки
	score := 0

	// Баллы за баланс
	if balance > 50000 {
		score += 40
	} else if balance > 20000 {
		score += 30
	} else if balance > 5000 {
		score += 20
	} else {
		score += 10
	}

	// Баллы за соотношение доход/расход
	if ratio > 1.5 {
		score += 40
	} else if ratio > 1.2 {
		score += 30
	} else if ratio > 0.8 {
		score += 20
	} else {
		score += 10
	}

	// Баллы за стабильность (на основе трендов)
	// (В реальной реализации нужно анализировать monthlyTrends)
	score += 20

	return min(score, 100)
}

package model

type CategorySpending struct {
	CategoryName     string  `json:"category_name"`
	TotalAmount      float64 `json:"total_amount"`
	TransactionCount int     `json:"transaction_count"`
	Percentage       float64 `json:"percentage,omitempty"` // Будет вычислено в сервисе
}

type FinancialHealth struct {
	TotalBalance       float64               `json:"total_balance"`
	IncomeExpenseRatio float64               `json:"income_expense_ratio"`
	LargestExpenseCat  string                `json:"largest_expense_category"`
	MonthlyTrends      []MonthlyBalanceTrend `json:"monthly_trends"`
	HealthScore        int                   `json:"health_score"`
}

type MonthlyBalanceTrend struct {
	Month   string  `json:"month"` // Формат: "Январь 2024"
	Balance float64 `json:"balance"`
}

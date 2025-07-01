package repository

import (
	"fin_tracker/internal/model"
	"time"

	"gorm.io/gorm"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) GetSpendingByCategory(userID string, start, end time.Time) ([]model.CategorySpending, error) {
	var result []model.CategorySpending

	err := r.db.Table("transactions").
		Select(`
            categories.name AS category_name,
            SUM(transactions.amount) AS total_amount,
            COUNT(transactions.id) AS transaction_count
        `).
		Joins("LEFT JOIN categories ON categories.id = transactions.category_id").
		Where("transactions.user_id = ?", userID).
		Where("transactions.transaction_date BETWEEN ? AND ?", start, end).
		Where("transactions.amount < 0"). // Только расходы
		Group("categories.name").
		Order("total_amount ASC").
		Scan(&result).Error

	return result, err
}

func (r *TransactionRepository) GetFinancialHealthData(userID string) (
	totalBalance float64,
	income float64,
	expenses float64,
	largestCategory string,
	monthlyTrends []model.MonthlyBalanceTrend,
	err error,
) {
	// Общий баланс
	r.db.Model(&model.Account{}).
		Select("COALESCE(SUM(initial_balance), 0)").
		Where("user_id = ?", userID).
		Scan(&totalBalance)

	// Доходы и расходы за последние 6 месяцев
	now := time.Now()
	sixMonthsAgo := now.AddDate(0, -6, 0)

	var transactions []struct {
		Amount float64
		Type   string
	}

	r.db.Table("transactions").
		Select("SUM(amount) as amount, CASE WHEN amount > 0 THEN 'income' ELSE 'expense' END as type").
		Where("user_id = ? AND transaction_date BETWEEN ? AND ?", userID, sixMonthsAgo, now).
		Group("type").
		Scan(&transactions)

	for _, t := range transactions {
		if t.Type == "income" {
			income = t.Amount
		} else {
			expenses = t.Amount
		}
	}

	// Крупнейшая категория расходов
	r.db.Table("transactions").
		Select("categories.name, SUM(transactions.amount) as total").
		Joins("JOIN categories ON categories.id = transactions.category_id").
		Where("transactions.user_id = ? AND transactions.amount < 0", userID).
		Group("categories.name").
		Order("total ASC"). // Отрицательные суммы, ASC = наименьшее значение (наибольший расход)
		Limit(1).
		Scan(&largestCategory)

	// Тренды по месяцам
	r.db.Raw(`
        SELECT 
            TO_CHAR(date_trunc('month', transaction_date), 'Month YYYY') AS month,
            SUM(amount) AS balance
        FROM transactions
        WHERE user_id = ? AND transaction_date >= ?
        GROUP BY date_trunc('month', transaction_date)
        ORDER BY date_trunc('month', transaction_date) ASC
    `, userID, sixMonthsAgo).Scan(&monthlyTrends)

	return totalBalance, income, expenses, largestCategory, monthlyTrends, nil
}

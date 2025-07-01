package router

import (
	_ "fin_tracker/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	"fin_tracker/internal/handler"
	"fin_tracker/internal/middleware"
	"fin_tracker/internal/repository"
	"fin_tracker/internal/service"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Репозитории
	userRepo := repository.NewUserRepository(db)
	accountRepo := repository.NewAccountRepository(db)
	currencyRepo := repository.NewCurrencyRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)
	healthService := service.NewFinancialHealthService(transactionRepo)

	// Сервисы
	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)
	accountService := service.NewAccountService(accountRepo, currencyRepo)
	analyticsService := service.NewAnalyticsService(transactionRepo)
	healthHandler := handler.NewFinancialHealthHandler(healthService)

	// Хендлеры
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	accountHandler := handler.NewAccountHandler(accountService)
	analyticsHandler := handler.NewAnalyticsHandler(analyticsService)

	// Публичные маршруты
	r.POST("/register", authHandler.Register)
	r.POST("/login", authHandler.Login)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Приватные маршруты с JWT middleware
	auth := r.Group("/")
	auth.Use(middleware.JWTAuthMiddleware())
	auth.GET("/me", userHandler.GetMe)
	auth.GET("/accounts", accountHandler.GetAccounts)
	auth.GET("/accounts/:id", accountHandler.GetAccount)
	auth.POST("/accounts", accountHandler.CreateAccount)
	auth.PATCH("/accounts/:id", accountHandler.UpdateAccount)
	auth.DELETE("/accounts/:id", accountHandler.DeleteAccount)
	auth.GET("/analytics/spending-by-category", analyticsHandler.GetSpendingByCategory)
	auth.GET("/analytics/financial-health", healthHandler.GetFinancialHealth)

	return r
}

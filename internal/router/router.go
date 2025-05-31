package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"fin_tracker/internal/handler"
	// "fin_tracker/internal/middleware"
	"fin_tracker/internal/repository"
	"fin_tracker/internal/service"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Репозитории
	userRepo := repository.NewUserRepository(db)

	// Сервисы
	// authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)

	// Хендлеры
	// authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)

	// Публичные маршруты
	// r.POST("/register", authHandler.Register)
	// r.POST("/login", authHandler.Login)

	// Приватные маршруты с JWT middleware
	auth := r.Group("/")
	// auth.Use(middleware.JWTAuthMiddleware())
	auth.GET("/me", userHandler.GetMe)

	return r
}

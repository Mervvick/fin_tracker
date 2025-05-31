package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"fin_tracker/internal/handler"
	"fin_tracker/internal/repository"
	"fin_tracker/internal/service"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Репозитории
	userRepo := repository.NewUserRepository(db)

	// Сервисы
	userService := service.NewUserService(userRepo)

	// Хендлеры
	userHandler := handler.NewUserHandler(userService)

	auth := r.Group("/")
	auth.GET("/me", userHandler.GetMe)

	return r
}

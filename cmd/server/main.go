// @title API fin_tracker
// @version 1.0
// @description ПР1-6.
// @host localhost:444
// @BasePath /
// @schemes https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

package main

import (
	"fin_tracker/internal/config"
	"fin_tracker/internal/model"
	"fin_tracker/internal/redis"
	"fin_tracker/internal/router"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/sirupsen/logrus"
)

func main() {
	redis.InitRedis()
	logFile, err := os.OpenFile("/app/logs/api.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logrus.SetOutput(logFile)
	} else {
		logrus.Warn("Не удалось создать файл логов, используем stdout")
	}

	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	err = db.AutoMigrate(
		&model.User{},
		&model.Currency{},
		&model.Account{},
		&model.Category{},
		&model.Transaction{},
		&model.RecurringTransaction{},
	)
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}

	log.Println("Подключение к БД успешно, миграция завершена")

	r := router.SetupRouter(db)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}

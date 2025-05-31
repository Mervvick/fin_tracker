package main

import (
	"fin_tracker/internal/config"
	"fin_tracker/internal/models"
	"fin_tracker/internal/router"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB connection errors: %v", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Account{},
		&models.Category{},
		&models.Operation{},
	)
	if err != nil {
		log.Fatalf("Migrations errors: %v", err)
	}

	log.Println("Connected to db successfully. Migration completed")

	r := router.SetupRouter(db)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server launch errors: %v", err)
	}
}

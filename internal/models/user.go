package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email        string    `gorm:"uniqueIndex;not null"`
	PasswordHash string    `gorm:"not null"`
	FullName     string
	CreatedAt    time.Time `gorm:"autoCreateTime"`

	Accounts   []Account   `gorm:"foreignKey:UserID"`
	Categories []Category  `gorm:"foreignKey:UserID"`
	Operation  []Operation `gorm:"foreignKey:UserID"`
}

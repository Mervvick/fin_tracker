package models

import (
	"time"

	"github.com/google/uuid"
)

type Operation struct {
	ID            uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID        uuid.UUID  `gorm:"type:uuid;not null"`
	AccountID     uuid.UUID  `gorm:"type:uuid;not null"`
	CategoryID    *uuid.UUID `gorm:"type:uuid"`
	Amount        float64    `gorm:"not null"`
	OperationName string
	Descroption   string
	Date          time.Time `gorm:"not null"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`

	Account  Account
	Category *Category
}

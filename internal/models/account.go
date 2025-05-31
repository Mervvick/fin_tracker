package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Name      string    `gorm:"not null"`
	Balance   float64
	CreatedAt time.Time `gorm:"autoCreateTime"`

	Operation []Operation `gorm:"foreignKey:AccountID"`
}

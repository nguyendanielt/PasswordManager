package model

import (
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	UserId       uuid.UUID `gorm:"primaryKey;not null"`
	ActivityType string    `gorm:"not null"`
	DateAndTime  time.Time `gorm:"primaryKey;not null"`
}

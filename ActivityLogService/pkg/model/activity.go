package model

import (
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	ActivityId   uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserId       uuid.UUID `gorm:"not null"`
	ActivityType string    `gorm:"not null"`
	DateAndTime  time.Time `gorm:"not null"`
}

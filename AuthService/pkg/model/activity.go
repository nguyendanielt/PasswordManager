package model

import (
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	UserId       uuid.UUID
	ActivityType string
	DateAndTime  time.Time
}

package model

import (
	"time"
)

type User struct {
	ID        int    `gorm:"primaryKey; autoIncrement"`
	Email     string `gorm:"unique; not null"`
	Username  string `gorm:"unique; not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
}

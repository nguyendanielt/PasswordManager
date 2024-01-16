package model

import (
	"github.com/google/uuid"
)

type Password struct {
	UserId      uuid.UUID `gorm:"primaryKey;not null"`
	AccountName string    `gorm:"primaryKey;not null"`
	Email       string
	Username    string
	Password    string `gorm:"not null"`
}

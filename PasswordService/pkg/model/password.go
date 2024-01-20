package model

import (
	"github.com/google/uuid"
)

type Password struct {
	PasswordId  uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserId      uuid.UUID `gorm:"not null"`
	AccountName string    `gorm:"not null"`
	Email       string
	Username    string
	Password    string `gorm:"not null"`
}

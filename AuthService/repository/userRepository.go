package repository

import (
	"errors"
	"fmt"

	"authservice/model"
	"authservice/util"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByEmailAndPassword(email string, password string) *model.User {
	var user model.User
	result := r.db.Model(&model.User{}).Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) || !util.CompareHashAndPwdString(user.Password, password) {
		return nil
	}
	return &user
}

func (r *UserRepository) AddUser(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		fmt.Println("Error when adding user:", err)
		return err
	}
	fmt.Println("Successfully added user")
	return nil
}

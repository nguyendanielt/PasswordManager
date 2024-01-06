package service

import (
	"userauthservice/model"
	"userauthservice/repository"
	"userauthservice/util"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(user *model.User) {
	user.Password = util.GenerateHashedPwdString(user.Password)
	s.userRepo.AddUser(user)
}

func (s *UserService) GetUser(user *model.User) *model.User {
	return s.userRepo.GetUserByEmailAndPassword(user.Email, user.Password)
}

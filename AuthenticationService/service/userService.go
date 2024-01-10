package service

import (
	"authenticationservice/model"
	"authenticationservice/repository"
	"authenticationservice/util"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(user *model.User) error {
	user.Password = util.GenerateHashedPwdString(user.Password)
	err := s.userRepo.AddUser(user)
	return err
}

func (s *UserService) GetUser(user *model.User) *model.User {
	return s.userRepo.GetUserByEmailAndPassword(user.Email, user.Password)
}

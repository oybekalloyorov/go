package service

import (
	"oybekalloyorov/salom/dostonbek/instagram/internal/models"
	"oybekalloyorov/salom/dostonbek/instagram/internal/repository"
)

type UserService struct {
	repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo)*UserService{
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) CreateUser(req *models.UserModel) (*models.UserModel, error){
	return s.repo.CreateUser(req)
}

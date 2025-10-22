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

func (s *UserService) GetUserById(id int)(*models.UserModel, error) {
	return s.repo.GetUserById(id)
}

func (s *UserService)GetAllUsers()([]*models.UserModel, error){
	return s.repo.GetAllUsers()
}

func (s *UserService)UpdateUserById(obj *models.UserModel)(*models.UserModel, error){
	response, err := s.repo.UpdateUserByID(obj)
	if err != nil {
		return nil, err
	}
	return  response, err
}

func (s *UserService) DeleteUserByID(id int) error{
	return s.repo.DeleteUserByID(id)
}
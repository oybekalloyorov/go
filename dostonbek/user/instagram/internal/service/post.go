package service

import (
	"instagram/internal/models"
	"instagram/internal/repository"
)

type PostService struct {
	repo *repository.PostRepo
}

func NewPostService(postRepo *repository.PostRepo) *PostService {
	return &PostService{
		repo: postRepo,
	}
}

func (s *PostService) UpdatePost(obj *models.Post) (*models.Post, error) {
	response, err := s.repo.UpdatePost(obj)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (s *PostService) CreatePost(req *models.Post) (*models.Post, error) {
	return s.repo.CreatePost(req)
}

func (s *PostService) GetAllPosts() ([]*models.Post, error) {
	return s.repo.GetAllPosts()
}

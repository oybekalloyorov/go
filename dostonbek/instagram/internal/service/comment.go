package service

import (
	"oybekalloyorov/salom/dostonbek/instagram/internal/models"
	"oybekalloyorov/salom/dostonbek/instagram/internal/repository"
)

type CommentService struct {
	repo *repository.CommentRepo
}

func NewCommentService(commentRepo *repository.CommentRepo) *CommentService{
	return &CommentService{
		repo: commentRepo,
	}
}

func (s *CommentService) UpdateComment(obj *models.Comment) (*models.Comment, error){
	response, err := s.repo.UpdateComment(obj)
	if err !=nil {
		return nil, err
	}

	return response, nil
}
func (s *CommentService) CreateComment(obj *models.Comment)(*models.Comment, error){
	return s.repo.CreateComment(obj)
}

func (s *CommentService) GetAllComments() ([]*models.Comment, error){
	return s.repo.GetAllComments()
}

func (s *CommentService) GetCommentById(id int)(*models.Comment, error){
	return s.repo.GetCommentById(id)
}

func (s *CommentService) GetCommentsByUserId(id int)([]*models.Comment, error){
	return s.repo.GetCommentsByUserId(id)
}


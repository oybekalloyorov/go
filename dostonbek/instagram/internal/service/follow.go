package service

import (
	"oybekalloyorov/salom/dostonbek/instagram/internal/models"
	"oybekalloyorov/salom/dostonbek/instagram/internal/repository"
)

type FollowService struct {
	FollowRepo *repository.FollowRepo
}

func NewFollowService(follovRepo *repository.FollowRepo) *FollowService{
	return &FollowService{
		FollowRepo: follovRepo,
	}
}

func (f *FollowService) CreateFollows(fol *models.Follow)(*models.Follow, error){
	return f.FollowRepo.CreateFollows(fol)
}
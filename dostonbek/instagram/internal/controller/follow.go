package controller

import (
	"log"
	"net/http"
	"oybekalloyorov/salom/dostonbek/instagram/internal/models"
	"oybekalloyorov/salom/dostonbek/instagram/internal/service"

	"github.com/gin-gonic/gin"
)

type FollowController struct {
	followService *service.FollowService
}

func NewFollowController(followService *service.FollowService) *FollowController{
	return &FollowController{
		followService: followService,
	}
}

func (f FollowController)CreateFollowHTTP(c *gin.Context){
	var req models.Follow

	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Println(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	res, err := f.followService.CreateFollows(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

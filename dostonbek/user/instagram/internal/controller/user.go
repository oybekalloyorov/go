package controller

import (
	"instagram/internal/models"
	"instagram/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	srv *service.UserService
}

func NewUserController(srv *service.UserService) *UserController {
	return &UserController{
		srv: srv,
	}
}

func (h *UserController) CreateUserHTTP(req *gin.Context) {
	var obj models.UserModel
	err := req.ShouldBindJSON(&obj)
	if err != nil {
		req.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	response, err := h.srv.CreateUser(&obj)
	if err != nil {
		req.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.JSON(http.StatusCreated, gin.H{"response": response})
}

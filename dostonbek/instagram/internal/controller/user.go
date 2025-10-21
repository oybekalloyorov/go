package controller

import (
	"net/http"
	"oybekalloyorov/salom/dostonbek/instagram/internal/models"
	"oybekalloyorov/salom/dostonbek/instagram/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	srv *service.UserService
}

func NewUserController(srv *service.UserService) *UserController{
	return &UserController{
		srv: srv,
	}
}

func (u *UserController) CreateUserHTTP(req *gin.Context){
	var obj models.UserModel
	err := req.ShouldBindJSON(&obj)
	if err != nil {
		req.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	response, err := u.srv.CreateUser(&obj)
	if err != nil {
		req.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.JSON(http.StatusCreated, gin.H{"response": response})
}

func (u *UserController) GetUserByIdHTTP(req *gin.Context){
	idStr := req.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		req.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	res, err := u.srv.GetUserById(id)
	if err != nil {
		req.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	req.JSON(http.StatusOK, res)
}

func (u *UserController)GetAllUsers(req *gin.Context){
	res, err := u.srv.GetAllUsers()
	if err != nil{
		req.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req.JSON(http.StatusOK, res)
}

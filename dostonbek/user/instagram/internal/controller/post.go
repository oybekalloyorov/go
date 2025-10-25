package controller

import (
	"instagram/internal/models"
	"instagram/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService *service.PostService
}

func NewPostController(postService *service.PostService) *PostController {
	return &PostController{
		postService: postService,
	}
}

func (p *PostController) UpdatePostHTTP(c *gin.Context) {
	var req models.Post
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	res, err := p.postService.UpdatePost(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": res})
}

func (p *PostController) CreatePostHTTP(c *gin.Context) {
	var req models.Post
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	res, err := p.postService.CreatePost(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": res})
}

func (p *PostController) GetAllPostsHTTP(c *gin.Context) {
	res, err := p.postService.GetAllPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": res})
}

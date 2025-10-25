package main

import (
	"instagram/internal/controller"
	"instagram/internal/pkg/config"
	"instagram/internal/pkg/helper"
	"instagram/internal/repository"
	"instagram/internal/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.DBConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "instagram",
		Password: "oybek",
		DBName:   "instagram",
		SSLMode:  "disable",
	}

	db, err := helper.ConnectPostgress(&cfg)
	if err != nil {
		panic(err)
	}

	// POSTS
	postRepo := repository.NewPostRepo(db)
	postService := service.NewPostService(postRepo)
	postController := controller.NewPostController(postService)

	// USERS
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	router := gin.Default()

	// post routes
	router.PATCH("/api/v1/update-post", postController.UpdatePostHTTP)
	router.POST("/api/v1/create-post", postController.CreatePostHTTP)
	router.GET("/api/v1/get-all-posts", postController.GetAllPostsHTTP)

	// user routes
	router.POST("/api/v1/create-user", userController.CreateUserHTTP)

	router.Run(":8000")
}

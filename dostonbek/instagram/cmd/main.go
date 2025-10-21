package main

import (
	"oybekalloyorov/salom/dostonbek/instagram/internal/controller"
	"oybekalloyorov/salom/dostonbek/instagram/internal/pkg/config"
	"oybekalloyorov/salom/dostonbek/instagram/internal/pkg/helper"
	"oybekalloyorov/salom/dostonbek/instagram/internal/repository"
	"oybekalloyorov/salom/dostonbek/instagram/internal/service"

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

	db, err := helper.ConnectPostgres(&cfg)
	if err != nil {
		panic(err)
	}

	// Posts
	postRepo := repository.NewPostRepo(db)
	postService := service.NewPostService(postRepo)
	postController := controller.NewPostController(postService)

	// Users
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	router := gin.Default()

	// Post Routes
	router.PATCH("/api/v1/update-post", postController.UpdatePostHTTP)
	router.POST("/api/v1/create-post", postController.CreatePostHTTP)
	router.GET("/api/v1/get-all-posts", postController.GetAllPostsHTTP)
	router.GET("/api/v1/get-posts-byid/:id", postController.GetToDoByIDHTTP)
	router.DELETE("/api/v1/delete/:id", postController.DeleteByIDHTTP)

	router.GET("/api/v1/get-post-by-userid/:id", postController.GetPostByUserIDHTTP)

	// User routes
	router.POST("/api/v1/create-user", userController.CreateUserHTTP)
	router.GET("/api/v1/get-user-byid/:id", userController.GetUserByIdHTTP)
	router.GET("/api/v1/get-all-users", userController.GetAllUsers)
	router.PATCH("/api/v1/update-user", userController.UpdateUserHTTP)

	router.Run(":8000")
}

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
	postRepo := repository.NewPostRepo(db)
	postService := service.NewPostService(postRepo)
	postController := controller.NewPostController(postService)

	router := gin.Default()

	router.PATCH("/api/v1/update-post", postController.UpdatePostHTTP)
	router.POST("/api/v1/create-post", postController.CreatePostHTTP)
	router.GET("/api/v1/get-all-posts", postController.GetAllPostsHTTP)

	router.Run(":8000")
}

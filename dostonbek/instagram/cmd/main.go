package main

import (
	"fmt"
	"net/http"
	"oybekalloyorov/salom/dostonbek/instagram/internal/controller"
	"oybekalloyorov/salom/dostonbek/instagram/internal/models"
	"oybekalloyorov/salom/dostonbek/instagram/internal/pkg/config"
	"oybekalloyorov/salom/dostonbek/instagram/internal/pkg/helper"
	"oybekalloyorov/salom/dostonbek/instagram/internal/repository"
	"oybekalloyorov/salom/dostonbek/instagram/internal/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"github.com/qor/admin"
	"github.com/qor/qor"
)
	type User struct {
		ID          int       `json:"id"`
		FullName    string    `json:"full_name"`
		Username    string    `json:"username"`
		BirthOfYear int       `json:"birth_of_year"`
		Bio         string    `json:"bio"`
		CreatedAt   time.Time `json:"created_at"`
	}

	func (User) TableName() string {
		return "instagram_users"
	}
	type Post models.Post
	type Comments models.Comment
	type Follows models.Follow
	func (Post) TableName() string{
		return "post"
	}
	func (Comments) TableName() string{
		return "comments"
	}
	func (Follows) TableName() string{
		return "follows"
	}

func main() {
	// cfg := config.DBConfig{
	// 	Host:     "localhost",
	// 	Port:     5432,
	// 	User:     "instagram",
	// 	Password: "oybek",
	// 	DBName:   "instagram",
	// 	SSLMode:  "disable",
	// }

	cfg := config.LoadConfig()

	db, err := helper.ConnectPostgres(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Database connection is ready ✅")
	// 
	
	// 🔹 GORM ulanish
	// dsn := "host=localhost user=instagram password=oybek dbname=instagram port=5432 sslmode=disable"
	gormDB, err := gorm.Open("postgres", db)
	// if err != nil {
	// 	panic(err)
	// }

	// 🔹 Model yaratish
	gormDB.AutoMigrate(&User{})

	// 🔹 QOR Admin panel
	Admin := admin.New(&qor.Config{DB: gormDB})
	Admin.AddResource(&User{}, &admin.Config{Menu: []string{"User"}})
	Admin.AddResource(&Post{}, &admin.Config{Menu: []string{"User Management"}})
	Admin.AddResource(&Comments{}, &admin.Config{Menu: []string{"User Management"}})
	Admin.AddResource(&Follows{}, &admin.Config{Menu: []string{"User Management"}})
	
	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)

	go func() {
		http.ListenAndServe(":9000", mux)
	}()
	// 
	// Posts
	postRepo := repository.NewPostRepo(db)
	postService := service.NewPostService(postRepo)
	postController := controller.NewPostController(postService)

	// Users
	userRepo := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	//Comments
	commentRepo := repository.NewCommentRepo(db)
	commentService := service.NewCommentService(commentRepo)
	commentController := controller.NewCommentController(commentService)

	//Follows
	followRepo := repository.NewFollowRepo(db)
	followService := service.NewFollowService(followRepo)
	followController := controller.NewFollowController(followService)


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
	router.DELETE("/api/v1/delete-user/:id", userController.DeleteUserHTTP)

	//Comments
	router.POST("/api/v1/create-comment", commentController.CreateCommentHTTP)
	router.GET("/api/v1/get-comment-by-id/:id", commentController.GetCommentByIdHTTP)
	router.GET("/api/v1/get-comment-by-user-id/:id", commentController.GetCommentsByUserIdHTTP)
	router.GET("/api/v1/get-all-comments", commentController.GetAllComments)
	router.PATCH("/api/v1/update-comment", commentController.UpdateCommentHTTP)
	
	//Folows
	router.POST("/api/v1/create-follow", followController.CreateFollowHTTP)
	
	router.Run(":8000")
}

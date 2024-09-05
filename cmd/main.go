package main

import (
	"database/sql"
	"github.com/Waelson/internal/controller"
	"github.com/Waelson/internal/repository"
	"github.com/Waelson/internal/service"
	"github.com/Waelson/internal/util/db"
	middleware2 "github.com/Waelson/internal/util/middleware"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	//Initialize database
	database, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	//Create all tables
	err = db.CreateTables(database)
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	//Custom Gin Validators
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("not_blank", middleware2.NotBlank)
		_ = v.RegisterValidation("tags", middleware2.ValidateTags)
		_ = v.RegisterValidation("duplicated", middleware2.ValidateDuplicatedTags)
	}

	//Repositories
	postRepository := repository.NewPostRepository(database)
	answerRepository := repository.NewAnswerRepository(database)
	userRepository := repository.NewUserRepository(database)
	//Services
	searchEngineService := service.NewSearchEngineService()
	notificationService := service.NewNotificationService()
	postService := service.NewPostService(postRepository, searchEngineService)
	answerService := service.NewAnswerService(answerRepository, postRepository, notificationService)
	userService := service.NewUserService(userRepository)
	//Controllers
	postController := controller.NewPostController(postService)
	answerController := controller.NewAnswerController(answerService)
	userController := controller.NewUserController(userService)

	//Opened routes
	r.POST("/api/v1/users/login", userController.Login)
	//Protected routes
	authorized := r.RouterGroup.Group("/", middleware2.Authentication())
	authorized.POST("/api/v1/posts", postController.Save)
	authorized.POST("/api/v1/answers", answerController.Save)
	authorized.POST("/api/v1/users", userController.Save)

	_ = r.Run()
}

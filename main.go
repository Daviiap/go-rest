package main

import (
	"go_rest_api/database"
	"go_rest_api/handlers"
	"go_rest_api/models"
	"go_rest_api/repositories"
	"go_rest_api/services"
	"go_rest_api/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func createUser(db *gorm.DB) {
	hashPasswd, _ := utils.HashPassword("123")

	user := &models.User{
		ID:       "1",
		Name:     "jon doe",
		Email:    "jondoe@email.com",
		Password: hashPasswd,
	}

	db.Create(user)
}

func main() {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		panic("error on connecting database")
	}

	userRepo := repositories.NewUserRepo(db)
	loginService := services.NewLoginService(userRepo)
	loginHandler := handlers.NewLoginHandler(loginService)

	createUser(db)

	router := gin.New()
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginHandler.HandleLogin)
	}

	router.Run("localhost:8080")
}

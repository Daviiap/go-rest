package main

import (
	"go_rest_api/database"
	"go_rest_api/handlers"
	"go_rest_api/repositories"
	"go_rest_api/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.GetDatabaseConnection()
	if err != nil {
		panic("error on connecting database")
	}

	userRepo := repositories.NewUserRepo(db)
	loginService := services.NewLoginService(userRepo)
	loginHandler := handlers.NewLoginHandler(loginService)

	router := gin.New()
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginHandler.HandleLogin)
	}

	router.Run("localhost:8080")
}

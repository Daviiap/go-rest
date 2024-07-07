package handlers

import (
	"go_rest_api/dtos"
	"go_rest_api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	loginService *services.LoginService
}

func NewLoginHandler(loginService *services.LoginService) *LoginHandler {
	return &LoginHandler{
		loginService: loginService,
	}
}

func (handler *LoginHandler) HandleLogin(c *gin.Context) {
	var userDTO dtos.LoginRequestDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	token, user, err := handler.loginService.Login(userDTO.Email, userDTO.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}

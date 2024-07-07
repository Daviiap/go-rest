package dtos

import "go_rest_api/models"

type LoginRequestDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponseDTO struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

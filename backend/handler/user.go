package handler

import (
	"latihan-assesment-radika/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUserHandler(c *gin.Context) {
	var inputUser entities.UserInput

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		return err.Error()
	}
	helper.
}

package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// using cursor to cerate new controller

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) GetUserById(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"id": "12",
	})
}

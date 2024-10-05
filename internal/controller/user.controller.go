package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-ecommerce-backend-api/internal/services"
)

// using cursor to cerate new controller

type UserController struct {
	userService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: services.MewUserService(),
	}
}

// controller -> service -> repo -> models -> dbs

func (uc *UserController) GetUserById(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"id": uc.userService.GetUserService(),
	})
}

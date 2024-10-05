package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ecommerce-backend-api/internal/controller"
	"github.com/go-ecommerce-backend-api/internal/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.AuthMiddleware())
	// group api version

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", controller.NewPongController().Pong)
		v1.GET("/user/1", controller.NewUserController().GetUserById)
	}
	return r
}

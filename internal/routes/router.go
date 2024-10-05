package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	// group api version

	v1 := r.Group("/v1")
	{
		v1.GET("/ping", Pong)
		v1.PUT("/ping", Pong)
		v1.POST("/ping", Pong)
		v1.PATCH("/ping", Pong)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/ping", Pong)
		v2.PUT("/ping", Pong)
		v2.POST("/ping", Pong)
		v2.PATCH("/ping", Pong)
	}

	return r
}

func Pong(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

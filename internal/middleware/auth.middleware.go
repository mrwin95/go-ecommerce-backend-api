package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ecommerce-backend-api/pkg/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("Authorization")
		if token != "valid-token" {
			response.ErrorResponse(ctx, response.ErrInvalidToken, "")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

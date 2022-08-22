package middlewares

import (
	"base-project-api/services"

	"github.com/gin-gonic/gin"
)

func AuthJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the client secret key
		authorization := ctx.GetHeader("Authorization")
		if authorization == "" {
			ctx.JSON(401, gin.H{
				"message": "Token not found",
			})
			ctx.AbortWithStatus(401)
			return
		}
		token := authorization[7:]

		if !services.NewJWTService().ValidateToken(token) {
			ctx.JSON(401, gin.H{
				"message": "Invalid token",
			})
			ctx.AbortWithStatus(401)
			return
		}
	}
}

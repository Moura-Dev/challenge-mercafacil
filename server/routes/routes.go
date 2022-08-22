package routes

import (
	"base-project-api/controllers"
	middlewares "base-project-api/server/middleware"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/")
	{
		main.POST("login", controllers.HandlerLogin)
		main.POST("/user", controllers.HandlerUser)
		routers := main.Group("/", middlewares.AuthJwt())

		{
			routers.POST("/", controllers.HandlerContact)

		}

	}
	return router
}

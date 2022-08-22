package routes

import (
	"base-project-api/controllers"
	middlewares "base-project-api/server/middleware"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/")
	{
		main.POST("login", controllers.Login)
		routers := main.Group("/", middlewares.AuthJwt())

		{
			routers.POST("/varejao", controllers.VarejaoController)
			routers.POST("/macapa", controllers.MacapaController)
			routers.POST("/header", controllers.HeaderController)
			routers.POST("/user", controllers.UserController)

		}

	}
	return router
}

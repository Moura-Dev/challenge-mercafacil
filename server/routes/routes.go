package routes

import (
	"base-project-api/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/")
	{
		routers := main.Group("/")
		{
			routers.POST("/varejao", controllers.VarejaoController)
			routers.POST("/macapa", controllers.MacapaController)

		}
	}
	return router
}

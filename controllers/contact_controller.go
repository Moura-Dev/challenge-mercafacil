package controllers

import (
	"base-project-api/models"
	"base-project-api/repository"
	"base-project-api/services"
	"fmt"

	"github.com/gin-gonic/gin"
)

func HandlerContact(ctx *gin.Context) {
	token := ctx.Request.Header.Get("authorization")
	token = token[7:]

	customer, err := services.NewJWTService().GetCustomerFromToken(token)
	if err != nil {
		fmt.Println(err)
	}
	switch customer {
	case "macapa":
		MacapaController(ctx)
	case "varejao":
		VarejaoController(ctx)
	}
}

func VarejaoController(ctx *gin.Context) {
	data := models.Contacts{}
	if ctx.ShouldBindJSON(&data) == nil {
		for _, value := range data.Infos {
			repository.CreateContactVarejao(ctx, &value)
		}
		ctx.JSON(201, gin.H{
			"message": "Contacts created successfully",
		})
	} else {
		ctx.JSON(400, gin.H{
			"message": "Bad Request",
		})
	}
}

func MacapaController(ctx *gin.Context) {
	data := models.Contacts{}
	if ctx.ShouldBindJSON(&data) == nil {
		for _, value := range data.Infos {
			repository.CreateContactMacapa(ctx, &value)
		}
		ctx.JSON(201, gin.H{
			"message": "Contacts created successfully",
		})
	} else {
		ctx.JSON(400, gin.H{
			"message": "Bad Request",
		})
	}
}

package controllers

import (
	"base-project-api/models"
	"base-project-api/repository"
	"base-project-api/services"

	"github.com/gin-gonic/gin"
)

func HandlerUser(ctx *gin.Context) {
	data := models.User{}
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Bad Request is invalid",
		})
		return
	}
	switch data.Customer {
	case "varejao":
		UserVarejao(ctx, &data)
	case "macapa":
		UserMacapa(ctx, &data)
	}
}

func UserVarejao(ctx *gin.Context, data *models.User) {
	data.Password = services.SHA256ENCODER(data.Password)

	err := repository.CreateUserPostgres(ctx, data)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "user already exist",
		})
		return

	}

	ctx.JSON(201, gin.H{
		"message": "User created successfully",
	})
}

func UserMacapa(ctx *gin.Context, data *models.User) {
	data.Password = services.SHA256ENCODER(data.Password)

	err := repository.CreateUserMySQL(ctx, data)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "user already exist",
		})
		return

	}

	ctx.JSON(201, gin.H{
		"message": "User created successfully",
	})
}

package controllers

import (
	"base-project-api/models"
	"base-project-api/repository"
	"base-project-api/services"

	"github.com/gin-gonic/gin"
)

func HandlerLogin(ctx *gin.Context) {
	data := models.Login{}

	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}
	switch data.Customer {
	case "varejao":
		LoginVarejao(ctx, &data)
	case "macapa":
		LoginMacapa(ctx, &data)

	default:
		ctx.JSON(400, gin.H{
			"error": "Customer Not Found",
		})

	}
}

func LoginVarejao(ctx *gin.Context, data *models.Login) {
	var user models.User
	// Get User in db

	user, err := repository.GetUserPostgres(ctx, data.Login)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	if user.Password != services.SHA256ENCODER(data.Password) {
		ctx.JSON(401, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(data.Customer)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"token": token,
	})
}

func LoginMacapa(ctx *gin.Context, data *models.Login) {

	var user models.User
	// Get User in db

	user, err := repository.GetUserMySQL(ctx, data.Login)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	if user.Password != services.SHA256ENCODER(data.Password) {
		ctx.JSON(401, gin.H{
			"error": "invalid credentials",
		})
		return
	}

	token, err := services.NewJWTService().GenerateToken(data.Customer)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"token": token,
	})
}

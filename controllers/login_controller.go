package controllers

import (
	"base-project-api/db"
	"base-project-api/models"
	"base-project-api/repository"
	"base-project-api/services"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
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
		Login(ctx, &data, db.ConnPostgres)
	case "macapa":
		Login(ctx, &data, db.ConnMysql)

	default:
		ctx.JSON(400, gin.H{
			"error": "Customer Not Found",
		})

	}
}

func Login(ctx *gin.Context, data *models.Login, database *sqlx.DB) {
	var user models.User
	// Get User in db

	user, err := repository.GetUser(ctx, data.Login, database)
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

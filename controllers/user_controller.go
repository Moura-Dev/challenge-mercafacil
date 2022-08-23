package controllers

import (
	"base-project-api/db"
	"base-project-api/models"
	"base-project-api/repository"
	"base-project-api/services"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
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
		CreateUser(ctx, &data, db.ConnPostgres)
	case "macapa":
		CreateUser(ctx, &data, db.ConnMysql)
	default:
	}

}

func CreateUser(ctx *gin.Context, data *models.User, database *sqlx.DB) {
	data.Password = services.SHA256ENCODER(data.Password)

	err := repository.CreateUser(ctx, data, database)
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

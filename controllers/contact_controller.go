package controllers

import (
	"base-project-api/db"
	"base-project-api/models"
	"base-project-api/repository"
	"base-project-api/services"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func HandlerContact(ctx *gin.Context) {
	token := ctx.Request.Header.Get("authorization")
	token = token[7:]

	customer, err := services.NewJWTService().GetCustomerFromToken(token)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch customer {
	case "macapa":
		CreateContact(ctx, db.ConnMysql, customer)
	case "varejao":
		CreateContact(ctx, db.ConnPostgres, customer)

	default:
		ctx.JSON(400, gin.H{
			"message": "Bad Request",
		})
		return
	}

}

func CreateContact(ctx *gin.Context, database *sqlx.DB, customer string) {
	data := models.Contacts{}
	if ctx.ShouldBindJSON(&data) == nil {
		for _, value := range data.Infos {
			repository.CreateContact(ctx, &value, database, customer)
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

package controllers

import (
	"base-project-api/models"
	"base-project-api/repository"
	"base-project-api/services"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func HeaderController(ctx *gin.Context) {

	customer := ctx.Request.Header.Get("customer")

	if customer == "varejao" {
		VarejaoController(ctx)
	} else if customer == "macapa" {

		MacapaController(ctx)
	} else {
		ctx.JSON(400, gin.H{
			"message": "Bad Request",
		})
	}
}

func VarejaoController(ctx *gin.Context) {
	data := models.Contacts{}
	if ctx.ShouldBindJSON(&data) == nil {
		for _, value := range data.Infos {
			repository.CreateContactVarejao(ctx, value.Nome, value.Celular)
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
	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		fmt.Println(err)
	}

	data := models.Contacts{}

	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err)
	}
	for _, value := range data.Infos {
		repository.CreateContactMacapa(ctx, value.Nome, value.Celular)
	}

	ctx.JSON(201, gin.H{
		"message": "Contacts created successfully",
	})

}

func UserController(ctx *gin.Context) {
	data := models.User{}
	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "Bad Request is invalid",
		})
	} else {

		data.Password = services.SHA256ENCODER(data.Password)

		//tratar erro de duplicidade

		repository.CreateUser(ctx, &data)

		ctx.JSON(201, gin.H{
			"message": "User created successfully",
		})
	}
}

func Login(ctx *gin.Context) {
	data := models.Login{}

	err := ctx.ShouldBindJSON(&data)
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	var user models.User
	// Get User in db
	user, err = repository.GetUser(ctx, data.Login)
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

	token, err := services.NewJWTService().GenerateToken()
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

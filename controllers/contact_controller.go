package controllers

import (
	"base-project-api/models"
	"base-project-api/repository"
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

package repository

import (
	database "base-project-api/db"
	"base-project-api/models"
	"base-project-api/utils"
	"context"
	"fmt"
	"strings"
)

// Create contact repository Postgres.
func CreateContactVarejao(ctx context.Context, nome string, celular string) {
	db := database.ConnPostgres
	celular = utils.MaskCellPhoneVarejo(celular)
	contactsMaps := []map[string]interface{}{
		{"Nome": nome, "Celular": celular},
	}

	_, err := db.NamedExec(`INSERT INTO contacts (nome, celular)
        VALUES (:Nome, :Celular)`, contactsMaps)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(contactsMaps)
	db.MustBegin().Commit()
}

// Create contact repository Mysql.
func CreateContactMacapa(ctx context.Context, nome string, celular string) {
	db := database.ConnMysql
	celular = utils.MaskCellPhoneMacapa(celular)
	contactsMaps := []map[string]interface{}{
		{"Nome": strings.ToUpper(nome), "Celular": celular},
	}

	_, err := db.NamedExec(`INSERT INTO contacts (nome, celular)
        VALUES (:Nome, :Celular)`, contactsMaps)
	if err != nil {
		fmt.Println(err)
	}

	db.MustBegin().Commit()
}

func CreateUser(ctx context.Context, user *models.User) {
	db := database.ConnMysql
	_, err := db.NamedExec(`INSERT INTO users (login, password)
		VALUES (:login, :password)`, user)
	if err != nil {
		fmt.Println(err)
	}
	db.MustBegin().Commit()
}

func GetUser(ctx context.Context, login string) (models.User, error) {
	db := database.ConnMysql
	user := models.User{}
	// create a interface for login 
	err := db.Get(&user, "SELECT * FROM users WHERE login = ?", login)
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	return user, nil

}

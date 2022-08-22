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
func CreateContactVarejao(ctx context.Context, contact *models.ContactInfo) {
	db := database.ConnPostgres
	celular := utils.MaskCellPhoneVarejo(contact.Celular)
	contactsMaps := []map[string]interface{}{
		{"Nome": contact.Nome, "Celular": celular},
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
func CreateContactMacapa(ctx context.Context, contact *models.ContactInfo) {
	db := database.ConnMysql
	celular := utils.MaskCellPhoneMacapa(contact.Celular)
	contactsMaps := []map[string]interface{}{
		{"Nome": strings.ToUpper(contact.Nome), "Celular": celular},
	}

	_, err := db.NamedExec(`INSERT INTO contacts (nome, celular)
        VALUES (:Nome, :Celular)`, contactsMaps)
	if err != nil {
		fmt.Println(err)
	}

	db.MustBegin().Commit()
}

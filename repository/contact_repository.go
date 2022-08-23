package repository

import (
	"base-project-api/models"
	"base-project-api/utils"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Create contact repository Postgres.
func CreateContact(ctx context.Context, contact *models.ContactInfo, database *sqlx.DB, customer string) {
	celular := utils.MaskPhone(contact.Celular, customer)
	contactsMaps := []map[string]interface{}{
		{"Nome": contact.Nome, "Celular": celular},
	}

	_, err := database.NamedExec(`INSERT INTO contacts (nome, celular)
        VALUES (:Nome, :Celular)`, contactsMaps)
	if err != nil {
		fmt.Println(err)
	}

	database.MustBegin().Commit()
}

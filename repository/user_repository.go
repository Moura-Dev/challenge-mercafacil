package repository

import (
	"base-project-api/models"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

func CreateUser(ctx context.Context, user *models.User, database *sqlx.DB) error {
	_, err := database.NamedExec(`INSERT INTO users (login, password, customer)
		VALUES (:login, :password, :customer)`, user)
	if err != nil {
		fmt.Println(err)
		return err
	}

	database.MustBegin().Commit()
	return nil
}

func GetUser(ctx context.Context, login string, dbConn *sqlx.DB) (models.User, error) {
	user := models.User{}
	// create a interface for login
	conditional := "WHERE login = "
	arg := map[string]interface{}{
		"login": login,
	}
	conditional += ":login"
	query := `SELECT * FROM users ` + conditional + `;`
	query, args, err := sqlx.Named(query, arg)
	if err != nil {
		return user, err
	}

	if query, args, err = sqlx.In(query, args...); err != nil {
		return user, err
	}
	query = dbConn.Rebind(query)
	row := dbConn.QueryRow(query, args...)
	row.Scan(&user.ID, &user.Login, &user.Password, &user.Customer)
	return user, nil

}

package repository

import (
	database "base-project-api/db"
	"base-project-api/models"
	"context"
	"fmt"
)

func CreateUserMySQL(ctx context.Context, user *models.User) error {
	db := database.ConnMysql
	_, err := db.NamedExec(`INSERT INTO users (login, password, customer)
		VALUES (:login, :password, :customer)`, user)
	if err != nil {
		fmt.Println(err)
		return err
	}

	db.MustBegin().Commit()
	return nil
}

func CreateUserPostgres(ctx context.Context, user *models.User) error {
	db := database.ConnPostgres
	_, err := db.NamedExec(`INSERT INTO users (login, password, customer)
		VALUES (:login, :password, :customer)`, user)
	if err != nil {
		fmt.Println(err)
		return err
	}

	db.MustBegin().Commit()
	return nil
}

func GetUserMySQL(ctx context.Context, login string) (models.User, error) {
	db := database.ConnMysql
	user := models.User{}
	// create a interface for login
	err := db.Get(&user, "SELECT * FROM users WHERE login =?", login)
	if err != nil && user.Login != "" {
		return user, err
	}

	return user, nil

}

func GetUserPostgres(ctx context.Context, login string) (models.User, error) {
	db := database.ConnPostgres
	user := models.User{}
	// create a interface for login
	err := db.Get(&user, "SELECT * FROM users WHERE login =$1", login)
	if err != nil && user.Login != "" {
		return user, err
	}

	return user, nil

}

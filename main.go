package main

import (
	database "base-project-api/db"
	"base-project-api/db/migrations"
	"base-project-api/server"
	"log"
)

func main() {
	db1, err := database.StartDBPostgres("localhost", "postgres", "postgres", "admin", "admin", "5432")
	if err != nil {
		log.Fatalln(err)
	}

	defer db1.Close()

	db2, err := database.StartDBMysql("localhost", "mysql", "admin", "root", "admin", "3306")
	if err != nil {
		log.Fatal(
			"Error connecting to the database: ", err)
	}
	defer db2.Close()
	db1.MustExec(migrations.Schema)
	db2.MustExec(migrations.Schema)

	s := server.NewServer()

	s.Run()
}

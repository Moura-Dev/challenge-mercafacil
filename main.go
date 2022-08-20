package main

import (
	database "base-project-api/db"
	"base-project-api/server"
	"log"
)

type User struct {
	Id                 int
	First, Last, Email string
}

func main() {
	Db1, err := database.StartDBPostgres("localhost", "postgres", "postgres", "admin", "admin", "5432")
	if err != nil {
		log.Fatalln(err)
	}

	defer Db1.Close()

	db2, err := database.StartDBMysql("localhost", "mysql", "admin", "root", "admin", "3306")
	if err != nil {
		log.Fatal(
			"Error connecting to the database: ", err)
	}
	defer db2.Close()

	s := server.NewServer()

	s.Run()
}

package database

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	err          error
	ConnPostgres *sqlx.DB
	ConnMysql    *sqlx.DB
)

func StartDBPostgres(localhost, Db, DbName, User, Password, Port string) (*sqlx.DB, error) {
	ctx := context.Background()
	ConnPostgres, err = sqlx.Open(Db, GetPostgresURI(localhost, Port, User, Password, DbName))
	if err != nil {
		fmt.Printf("Error connecting to the database: %s", err)
		return ConnPostgres, err
	}

	ConnPostgres.SetMaxIdleConns(100)
	ConnPostgres.SetMaxOpenConns(100)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err = ConnPostgres.PingContext(ctx)
	if err != nil {
		fmt.Printf("Error connecting to the database: %s", err)
		return ConnPostgres, err
	}

	fmt.Println("Database connected", DbName)
	return ConnPostgres, err
}

func StartDBMysql(localhost, Db, DbName, User, Password, Port string) (*sqlx.DB, error) {
	ctx := context.Background()
	ConnMysql, err = sqlx.Connect(Db, GetMysqlURI(localhost, Port, User, Password, DbName))

	ConnMysql.SetMaxIdleConns(100)
	ConnMysql.SetMaxOpenConns(100)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	err = ConnMysql.PingContext(ctx)

	if err != nil {
		fmt.Printf("Error connecting to the database: %s", err)
		return ConnMysql, err
	}

	fmt.Println("Database connected", DbName)
	return ConnMysql, err
}

func GetPostgresURI(host string, port string, username string, password string, dbName string) string {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable", host, port, username, password, dbName)
	return psqlInfo
}

func GetMysqlURI(host string, port string, username string, password string, dbName string) string {
	mysqlInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", username, password, host, port, dbName)
	return mysqlInfo
}

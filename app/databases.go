package app

import (
	"database/sql"
	"os"
	"restapi-bank-scraper/helper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database_name := os.Getenv("DB_DATABASE_NAME")

	database := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database_name + "?charset=utf8&parseTime=true&loc=Local"

	db, err := sql.Open("mysql", database)
	helper.PanicIfError(err)
	return db
}

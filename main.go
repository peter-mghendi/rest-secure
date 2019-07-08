package main

import (
	"fmt"
	"os"
	a "rest-secure/app"

	"github.com/joho/godotenv"
)

var app a.App
var dbHost, dbUser, dbName, dbPass, dbType, port string

func init() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	dbHost = os.Getenv("db_host")
	dbUser = os.Getenv("db_user")
	dbName = os.Getenv("db_name")
	dbPass = os.Getenv("db_pass")
	dbType = os.Getenv("db_type")

	port = os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
}

func main() {
	app := a.App{}
	app.Init(dbHost, dbUser, dbName, dbPass, dbType)
	app.Run(port)
}

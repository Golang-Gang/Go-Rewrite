package main

import (
	"os"

	"github.com/joho/godotenv"
	app "github.com/Golang-Gang/Go-Rewrite/goServer"
	setup "github.com/Golang-Gang/Go-Rewrite/goServer/setup"
)

func main() {
	godotenv.Load(".env")
	a := app.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
		os.Getenv("APP_DB_HOST"))

	setup.SetupTables(a.DB)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	a.Run(port)
}
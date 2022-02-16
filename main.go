package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))
		setupTables(a.DB)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	a.Run(port)
}

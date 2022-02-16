package gotest

import (
	"os"
	"os/exec"
	"log"

	"github.com/joho/godotenv"
	app "github.com/Golang-Gang/Go-Rewrite/goServer"
)

func main() {
	godotenv.Load(".env")
	a := app.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	app.SetupTables(a.DB)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	a.Run(port)

	cmd := exec.Command("jest --verbose --runInBand --testLocationInResults --setupFiles dotenv/config")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		os.Exit(1)
	}
}
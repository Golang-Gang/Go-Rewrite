package main

import (
	"log"
	"os"
	"os/exec"
	"testing"
	"time"

	app "github.com/Golang-Gang/Go-Rewrite/goServer"
	"github.com/joho/godotenv"
)

var a app.App

func TestMain(m *testing.M) {
    /*
	godotenv.Load(".env")
    a.Initialize(
        os.Getenv("APP_DB_USERNAME"),
        os.Getenv("APP_DB_PASSWORD"),
        os.Getenv("APP_DB_NAME"))

    app.SetupTables(a.DB)
    code := m.Run()
    os.Exit(code)
    */
    
	godotenv.Load(".env")
	a := app.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
        os.Getenv("APP_DB_HOST"))

	app.SetupTables(a.DB)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	go a.Run(port)

    time.Sleep(5 * time.Second)

	cmd := exec.Command("npm", "test")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
		os.Exit(1)
	}
    os.Exit(0)
}
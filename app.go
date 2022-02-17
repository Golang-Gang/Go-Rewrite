package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"net/http"

	dogsController "github.com/Golang-Gang/Go-Rewrite/goServer/controllers/dogs"
	catsController "github.com/Golang-Gang/Go-Rewrite/goServer/controllers/cats"
	productsController "github.com/Golang-Gang/Go-Rewrite/goServer/controllers/products"
	resetController "github.com/Golang-Gang/Go-Rewrite/goServer/controllers/reset"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	//get the heroku connection string hopefully
	envConString := os.Getenv("DATABASE_URL")
	if	envConString != "" {
		connectionString = envConString;
	}
	
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}

func (a *App) initializeRoutes() {
	s := a.Router.PathPrefix("/products").Subrouter()
	productsController.AddRoutes(s, a.DB)
	dogSubRouter := a.Router.PathPrefix("/dogs").Subrouter()
	dogsController.AddRoutes(dogSubRouter, a.DB)
	catSubRouter := a.Router.PathPrefix("/cats").Subrouter()
	catsController.AddRoutes(catSubRouter, a.DB)
	resetSubRouter := a.Router.PathPrefix("/reset").Subrouter()
	resetController.AddRoutes(resetSubRouter, a.DB)
}

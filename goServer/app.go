package goServer

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"bytes"

	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	catsController "github.com/Golang-Gang/Go-Rewrite/goServer/controllers/cats"
	dogsController "github.com/Golang-Gang/Go-Rewrite/goServer/controllers/dogs"
	productsController "github.com/Golang-Gang/Go-Rewrite/goServer/controllers/products"
	dbResetController "github.com/Golang-Gang/Go-Rewrite/goServer/controllers/reset"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func addSpaces(s string) string {
	buf := &bytes.Buffer{}
	for _, rune := range s {
			buf.WriteRune(rune)
			buf.WriteRune(' ')
	}
	return buf.String()
}

func (a *App) Initialize(user, password, dbname, host string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", user, password, dbname, host)

	//get the heroku connection string hopefully
	envConString := os.Getenv("DATABASE_URL")
	if	envConString != "" {
		connectionString = envConString;
	}
	connectionString = connectionString + " sslmode=" + os.Getenv("DB_SSL");
	log.Println("db con string: ");
	log.Println(addSpaces(connectionString));
	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(port string) {
	log.Fatal(http.ListenAndServe(":" + port, a.Router))
}

func (a *App) initializeRoutes() {
	s := a.Router.PathPrefix("/products").Subrouter()
	productsController.AddRoutes(s, a.DB)

	catSubRouter := a.Router.PathPrefix("/cats").Subrouter()
	catsController.AddRoutes(catSubRouter, a.DB)

	dogSubRouter := a.Router.PathPrefix("/dogs").Subrouter()
	dogsController.AddRoutes(dogSubRouter, a.DB)

	resetRouter := a.Router.PathPrefix("/reset").Subrouter()
	dbResetController.AddRoutes(resetRouter, a.DB)
}

package products

import (
	"database/sql"
	"fmt"

	"encoding/json"
	"net/http"
	"strconv"

	dog "github.com/Golang-Gang/Go-Rewrite/goServer/models/dog"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Products struct {
	Router *mux.Router
	DB     *sql.DB
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func AddRoutes(r *mux.Router, db *sql.DB) {
	//GET /:id
	r.HandleFunc("/{id:[0-9]+}", func (w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
				respondWithError(w, http.StatusBadRequest, "Invalid product ID")
				return
		}
	
		pm := dog.Dog{ID: id}
		if err := pm.GetDog(db); err != nil {
				switch err {
				case sql.ErrNoRows:
						respondWithError(w, http.StatusNotFound, "Product not found")
				default:
						respondWithError(w, http.StatusInternalServerError, err.Error())
				}
				return
		}
	
		respondWithJSON(w, http.StatusOK, pm)

	}).Methods("GET")


    
	// GET /
	r.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
    count, _ := strconv.Atoi(r.FormValue("count"))
    start, _ := strconv.Atoi(r.FormValue("start"))

    if count > 10 || count < 1 {
        count = 10
    }
    if start < 0 {
        start = 0
    }

    dogs, err := dog.GetDogs(db, start, count)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, dogs)
	}).Methods("GET")



	// DELETE /:id
	r.HandleFunc("/{id:[0-9]+}", func (w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])// converting from a string to a number
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid Doggie ID")
        return
    }

    d := dog.Dog{ID: id}
    if err := d.DeleteDog(db); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	}).Methods("DELETE")



	// POST /
	r.HandleFunc("", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("this ran")
    var d dog.Dog
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&d); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()

    if err := d.CreateDog(db); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusCreated, d)
	}).Methods("POST")

	// PUT /:id
	r.HandleFunc("/{id:[0-9]+}", func (w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid doggie ID")
        return
    }

    var d dog.Dog
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&d); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
        return
    }
    defer r.Body.Close()
    d.ID = id

    if err := d.UpdateDog(db); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, d)
	}).Methods("PUT")
}

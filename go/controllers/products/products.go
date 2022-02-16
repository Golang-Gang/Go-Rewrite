package products

import (
	"database/sql"
	"fmt"

	"encoding/json"
	"net/http"
	"strconv"

	product "github.com/Golang-Gang/Go-Rewrite/go/models/product"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Not necessary for application to work
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
	
		pm := product.Product{ID: id}
		if err := pm.GetProduct(db); err != nil {
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

    products, err := product.GetProducts(db, start, count)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, products)
	}).Methods("GET")

	// DELETE /:id
	r.HandleFunc("/{id:[0-9]+}", func (w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"]) // Converts from string to a number?
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid Product ID")
        return
    }

    p := product.Product{ID: id} // creates an instance of our product struct and appends an id. 
    if err := p.DeleteProduct(db); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	}).Methods("DELETE")

	// POST /
	r.HandleFunc("", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("this ran")
    var p product.Product
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&p); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()

    if err := p.CreateProduct(db); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusCreated, p)
	}).Methods("POST")

	// PUT /:id
	r.HandleFunc("/{id:[0-9]+}" /*<- Likely regex, ensures id is a series of numbers between 0 and 9.*/, func (w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid product ID")
        return
    }

    var p product.Product
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&p); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid resquest payload")
        return
    }
    defer r.Body.Close()
    p.ID = id

    if err := p.UpdateProduct(db); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, p)
	}).Methods("PUT")
}

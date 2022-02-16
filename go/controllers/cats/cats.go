// declare package
package cats

// import packages/modules
import (
	"database/sql"
	"fmt"

	"encoding/json"
	"net/http"
	"strconv"

	cat "github.com/Golang-Gang/Go-Rewrite/go/models/cat"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// accessible via Cats.respondWithError
// the camelcase capitalization indicates a private method, meaning it can only be accessed from within the cat struct
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
				respondWithError(w, http.StatusBadRequest, "Invalid cat ID")
				return
		}
	
		catModelInstance := cat.Cat{ID: id}
		if err := catModelInstance.GetCat(db); err != nil {
				switch err {
				case sql.ErrNoRows:
						respondWithError(w, http.StatusNotFound, "Cat not found")
				default:
						respondWithError(w, http.StatusInternalServerError, err.Error())
				}
				return
		}
	
		respondWithJSON(w, http.StatusOK, catModelInstance)

	}).Methods("GET")

	// GET /
	r.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {

    cats, err := cat.GetCats(db)
    if err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, cats)
	}).Methods("GET")

	// DELETE /:id
	r.HandleFunc("/{id:[0-9]+}", func (w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"]) // Converts from string to a number?
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid Cat ID")
        return
    }

    p := cat.Cat{ID: id} // creates an instance of our cat struct and appends an id. 
    if err := p.DeleteCat(db); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
	}).Methods("DELETE")

	// POST /
	r.HandleFunc("", func (w http.ResponseWriter, r *http.Request) {
		fmt.Println("this ran")
    var cat cat.Cat
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&cat); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()

    if err := cat.CreateCat(db); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusCreated, cat)
	}).Methods("POST")

	// PUT /:id
	r.HandleFunc("/{id:[0-9]+}" /*<- Likely regex, ensures id is a series of numbers between 0 and 9.*/, func (w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid Cat ID")
        return
    }

    var cat cat.Cat
    decoder := json.NewDecoder(r.Body)
    if err := decoder.Decode(&cat); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }
    defer r.Body.Close()
    cat.ID = id

    if err := cat.UpdateCat(db); err != nil {
        respondWithError(w, http.StatusInternalServerError, err.Error())
        return
    }

    respondWithJSON(w, http.StatusOK, cat)
	}).Methods("PUT")
}
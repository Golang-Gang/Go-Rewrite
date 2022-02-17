// declare package
package reset

// import packages/modules
import (
	"database/sql"
	"fmt"

	"encoding/json"
	"net/http"

	setup "github.com/Golang-Gang/Go-Rewrite/goServer/setup"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func AddRoutes(r *mux.Router, db *sql.DB) {
	// GET /
	r.HandleFunc("", func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("restting db")
		setup.SetupTables(db);
		respondWithJSON(w, http.StatusOK, []int{1, 2, 3})
	}).Methods("GET")
}

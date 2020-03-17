package routes

// imports omited
import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func MakeHTTPHandler() http.Handler {
	fmt.Println("hfgjsf")
	r := mux.NewRouter()
	r.HandleFunc("/register", RegisterUser).Methods("POST")
	// r.HandleFunc("GET", "/bar", list)
	// r.HandleFunc("GET", "/bar/:id", get)
	// r.HandleFunc("PATCH", "/bar/:id", update)
	// r.HandleFunc("DELETE", "/bar/:id", delete)
	return r
}

package internal

import (
	"net/http"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/register", RegisterHandler).Methods("POST")
	r.HandleFunc("/login", LoginHandler).Methods("POST")
}

func RegisterHandler(w http.responseWriter, r *http.Request) {
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
}

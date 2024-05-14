package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/tylermac92/go-chat-app/internal"
)

func main() {
	r := mux.NewRouter()
	internal.RegisterRoutes(r)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r) err != nil {
		log.Fatalf("could not start server: %v", err)
	}
}

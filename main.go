package main

import (
	"log"
	"mux/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/api/v1/CreateBooks", handlers.CreateBooks).Methods("POST")

	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}

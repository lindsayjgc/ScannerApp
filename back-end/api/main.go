package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func InitializeRouter() {
	r := mux.NewRouter();

	// Subrouter for handling all requests made to API URL
	s := r.PathPrefix("/api").Subrouter();

	s.HandleFunc("/login", Login).Methods("POST")
	s.HandleFunc("/logged-in", IsLoggedIn).Methods("GET")

	// Handle CORS options
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
		AllowedMethods: []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"},
	})
	handler := c.Handler(r)

	log.Println("Listening on port 9000...");
	log.Fatal(http.ListenAndServe(":9000", handler))
}

func main() {
	InitialUserMigration()
	InitializeRouter()
}
package main

import (
	"log"
	"net/http"
	"os"

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

	listenMsg := "Listening on port " + os.Getenv("PORT") + "..."
	log.Println(listenMsg);
	log.Fatal(http.ListenAndServe(":" + os.Getenv("PORT"), handler))
}

func main() {
	InitialUserMigration()
	InitializeRouter()
}
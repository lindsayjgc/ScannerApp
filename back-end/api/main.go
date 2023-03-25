package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var r = mux.NewRouter()
var err error

const DB_PATH = "../db/groceryapp.db"

var jwtKey []byte

func InitializeRouter() {
	// Subrouter for handling all requests made to API URL
	s := r.PathPrefix("/api").Subrouter()

	s.HandleFunc("/signup", SignUp).Methods("POST")
	s.HandleFunc("/login", Login).Methods("POST")
	s.HandleFunc("/logout", Logout).Methods("POST")
	s.HandleFunc("/delete-user", DeleteUser).Methods("DELETE")
	s.HandleFunc("/logged-in", IsLoggedIn).Methods("GET")
	s.HandleFunc("/user-info", UserInfo).Methods("GET")
	s.HandleFunc("/add-allergies", AddAllergy).Methods("PUT")
	s.HandleFunc("/delete-allergies", DeleteAllergy).Methods("DELETE")
	s.HandleFunc("/check-allergies", CheckAllergies).Methods("POST")
	s.HandleFunc("/create-list", CreateList).Methods("POST")
	s.HandleFunc("/add-list-items", AddGroceryItem).Methods("POST")
	s.HandleFunc("/delete-lists", DeleteList).Methods("DELETE")

}

func StartServer() {
	// Handle CORS options
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"},
	})
	handler := c.Handler(r)

	// create a logging middleware that wraps the router
	loggingMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			next.ServeHTTP(w, r)
			duration := time.Now().Sub(startTime)
			log.Printf("[%s] %s %s (%s)", r.Method, r.URL.Path, r.RemoteAddr, duration)
		})
	}

	// register handler to logger
	loggedRouter := loggingMiddleware(handler)

	listenMsg := "Listening on port " + os.Getenv("PORT") + "..."
	log.Println(listenMsg)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), loggedRouter))
}

func main() {
	InitialUserMigration()
	InitialAllergyMigration()
	InitialListMigration()
	InitializeRouter()
	StartServer()
}

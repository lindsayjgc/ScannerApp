package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var CodeDB *gorm.DB

type Email struct {
	Email string `json:"email"`
}

type RawCode struct {
	Email string `json:"email"`
	Code string `json:"code"`
}

func InitialCodeMigration() {
	// var err error
	CodeDB, err = gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Error connecting to CodeDB.")
	}

	err = godotenv.Load()

	if err != nil {
		fmt.Println(err)
		panic("Error loading .env file.")
	}

	// AutoMigrate checks the DB for a matching existing schema - if it does
	// not exist, create/update the new schema
	CodeDB.AutoMigrate(&Code{})
}


func VerifyEmailSignup(w http.ResponseWriter, r *http.Request) {
	IssueCode(w, r, "signup")
}

func IssueCode(w http.ResponseWriter, r *http.Request, emailType string) {
	w.Header().Set("Content-Type", "application/json")

	// Decode email from frontend
	var email Email
	err = json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON body"))
		return
	}

	randomCode := GenerateRandomCode()
	err = SendEmail(email.Email, randomCode, "signup")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenerateResponse("Verification email sent successfully"))
}


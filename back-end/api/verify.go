package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var CodeDB *gorm.DB

type Code struct {
	gorm.Model
	Email string 
	Code string
	Expiration time.Time `gorm:"type:datetime"`
}

type Email struct {
	Email string `json:"email"`
}

type RawCode struct {
	Email string `json:"email"`
	Code string `json:"code"`
}

type ResetData struct {
	Email string `json:"email"`
	Password string `json:"password"`
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

func VerifyEmailReset(w http.ResponseWriter, r *http.Request) {
	IssueCode(w, r, "reset")
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
	err = SendEmail(email.Email, randomCode, emailType)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	// Delete any existing codes
	var deletedCode Code
	CodeDB.Where("email = ?", email.Email).Delete(&deletedCode)

	// Create expiration time for code
	expiration := time.Now().Add(5* time.Minute)

	// Create code object and save to DB for future checking
	code := Code{
		Email: email.Email,
		Code: randomCode,
		Expiration: expiration,
	}
	CodeDB.Create(&code)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenerateResponse("Verification email sent successfully"))
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var resetData ResetData
	err = json.NewDecoder(r.Body).Decode(&resetData)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON body"))
		return
	}

	// Ensure all fields are not empty
	v := reflect.ValueOf(resetData)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(GenerateResponse("All fields are required."))
			return
		}
	}

	// Check if email already exists
	var userSearch User
	result := UserDB.First(&userSearch, "email = ?", resetData.Email)
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GenerateResponse("Email not found"))
		return
	}

	// If email does not exist, encrypt password for storage
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(resetData.Password), 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Could not generate password hash"))
		return
	}

	userSearch.Password = string(passwordHash)
	UserDB.Save(&userSearch)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenerateResponse("Password reset successfully"))
}
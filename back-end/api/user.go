package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"reflect"

	"github.com/glebarez/sqlite"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var UserDB *gorm.DB

type User struct {
	gorm.Model        // Declare this as the schema for GORM
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func InitialUserMigration() {
	UserDB, err = gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Error connecting to UserDB.")
	}

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
		panic("Error loading .env file.")
	}

	jwtKey = []byte(os.Getenv("SECRET_KEY"))

	// AutoMigrate checks the DB for a matching existing schema - if it does
	// not exist, create/update the new schema
	UserDB.AutoMigrate(&User{})
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
	err = json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON body"))
		return
	}

	// Ensure all fields are not empty
	v := reflect.ValueOf(user)
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).Interface() == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(GenerateResponse("All fields are required."))
			return
		}
	}

	// Check if email already exists
	var checkEmail User
	result := UserDB.First(&checkEmail, "email = ?", user.Email)

	if result.RowsAffected != 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GenerateResponse("Email is already registered to an account"))
		return
	}

	// If email does not exist, encrypt password for storage
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Could not generate password hash"))
		return
	}
	// Create credentials for use later in creating cokoie
	credentials := Credentials{Email: user.Email, Password: user.Password}
	
	user.Password = string(passwordHash)
	UserDB.Create(&user)

	// Now that user is created, log them in
	err, statusCode := CreateCookie(w, credentials)

	if err != nil {
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(GenerateResponse("User successfully created"))
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Store credentials sent by client
	var credentials Credentials
	err = json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON body"))
		return
	}

	// Fetch the user record with email matching the email passed in
	var user User
	result := UserDB.First(&user, "email = ?", credentials.Email)

	// Handle email not connected to any user in DB
	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GenerateResponse("Email not registered to any account"))
		return
	}

	// Compare password credentials to stored hashed password
	bycrptErr := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(credentials.Password))

	if bycrptErr != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(GenerateResponse("Incorrect password"))
		return
	}

	// Process for creating a cookie to store logged in user email
	err, statusCode := CreateCookie(w, credentials)

	if err != nil {
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
	}

	w.WriteHeader(http.StatusOK);
	json.NewEncoder(w).Encode(GenerateResponse("User successfully logged in"))
}

func IsLoggedIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for an existing cookie and handle possible errors
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK);
	res := make(map[string]string)
	res["message"] = "User is currently logged in"
	res["email"] = claims.Email
	json.NewEncoder(w).Encode(res)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	// If cookie is obtained without errors, delete it and respond
	DeleteCookie(w)

	w.WriteHeader(http.StatusOK);
	res := make(map[string]string)
	res["message"] = "User successfully logged out"
	res["email"] = claims.Email;
	json.NewEncoder(w).Encode(res)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	UserDB.Where("email LIKE ?", claims.Email).Delete(&User{})
	AllergyDB.Where("email LIKE ?", claims.Email).Delete(&Allergy{})
	DeleteCookie(w)

	w.WriteHeader(http.StatusOK)
	res := make(map[string]string)
	res["message"] = "User successfully deleted"
	res["email"] = claims.Email;
	json.NewEncoder(w).Encode(res)
}
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
const DB_PATH = "../db/groceryapp.db"
var jwtKey = []byte("secret_key")

type User struct {
	gorm.Model // Declare this as the schema for GORM
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type Credentials struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func InitialUserMigration() {
	DB, err = gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Error connecting to DB.")
	}

	// AutoMigrate checks the DB for a matching existing schema - if it does
	// not exist, create/update the new schema
	DB.AutoMigrate(&User{})
}

func Login(w http.ResponseWriter, r *http.Request) {

	// Store credentials sent by client
	var credentials Credentials
	err := json.NewDecoder(r.Body).Decode(&credentials)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Fetch the user record with email matching the email passed in
	var user User
	result := DB.First(&user, "email = ?", credentials.Email)

	// Handle email not connected to any user in DB
	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Email not registered to any account."))
		return
	}

	if credentials.Password != user.Password {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Error: Incorect password."))
		return
	}

	expirationTime := time.Now().Add(time.Hour * 24) // JWT lasts 1 day
	claims := &Claims {
		Email: credentials.Email,
		StandardClaims: jwt.StandardClaims {
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating JWT."))
		return
	}

	http.SetCookie(w, 
	&http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: expirationTime,
		HttpOnly: true,
	})
}

func Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s", claims.Email)))
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	
}
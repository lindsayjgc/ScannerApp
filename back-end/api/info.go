package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Info struct {
	gorm.Model
	Email     string `json:"email"`
	Allergies string `json:"allergies"`
}

func InitialInfoMigration() {
	DB, err = gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Error connecting to DB.")
	}

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
		panic("Error loading .env file.")
	}

	jwtKey = []byte(os.Getenv("SECRET_KEY"))

	// AutoMigrate checks the DB for a matching existing schema - if it does
	// not exist, create/update the new schema
	DB.AutoMigrate(&Info{})
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	var email Email
	err := json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := make(map[string]string)
		res["msg"] = "Email Not Found"
		json.NewEncoder(w).Encode(res)
		return
	}

	var user User
	result := DB.First(&user, "email = ?", email.Email)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		res := make(map[string]string)
		res["msg"] = "User Not Found"
		json.NewEncoder(w).Encode(res)
		return
	}

	var info Info
	result = DB.First(&info, "email = ?", email.Email)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		res := make(map[string]string)
		res["msg"] = "User Info Not Found"
		json.NewEncoder(w).Encode(res)
		return
	}

	json.NewEncoder(w).Encode(user)
	json.NewEncoder(w).Encode(info)
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Info struct {
	gorm.Model
	Email     string `json:"email"`
	Allergies string `json:"allergies"`
}

type AllUserInfo struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Allergies string `json:"allergies"`
}

type Email struct {
	Email string `json:"email"`
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

	// all important user info combined into one struct for easier use by frontend
	var allInfo AllUserInfo
	allInfo.FirstName = user.FirstName
	allInfo.LastName = user.LastName
	allInfo.Email = email.Email
	allInfo.Password = user.Password
	allInfo.Allergies = info.Allergies

	json.NewEncoder(w).Encode(allInfo)
}

func AddAllergy(w http.ResponseWriter, r *http.Request) {
	var info Info

	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		res := make(map[string]string)
		res["msg"] = "Cannot decode user info"
		json.NewEncoder(w).Encode(res)
		return
	}

	var userInfo Info
	result := DB.First(&userInfo, "email = ?", info.Email)
	// if info is not found, create an entry for the user
	if result.Error != nil {
		DB.Create(&info)
		return
	}

	// check if user already has allergy logged, if not, add it
	allergies := userInfo.Allergies
	allergyList := strings.Split(allergies, ",")
	for i := 0; i < len(allergyList); i++ {
		if info.Allergies == allergyList[i] {
			res := make(map[string]string)
			res["msg"] = "Allergy already added"
			json.NewEncoder(w).Encode(res)
			return
		}
	}
	userInfo.Allergies += "," + info.Allergies
	DB.Save(&userInfo)
}

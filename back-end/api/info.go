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

var InfoDB *gorm.DB

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
	InfoDB, err = gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

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
	InfoDB.AutoMigrate(&Info{})
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var email Email
	err := json.NewDecoder(r.Body).Decode(&email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON"))
		return
	}

	var user User
	result := InfoDB.First(&user, "email = ?", email.Email)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(GenerateResponse("User not found"))
		return
	}

	var info Info
	var allergies = true
	result = InfoDB.First(&info, "email = ?", email.Email)
	if result.Error != nil {
		allergies = false
	}

	// all important user info combined into one struct for easier use by frontend
	var allInfo AllUserInfo
	allInfo.FirstName = user.FirstName
	allInfo.LastName = user.LastName
	allInfo.Email = email.Email
	allInfo.Password = user.Password
	if allergies == true {
		allInfo.Allergies = info.Allergies
	} else {
		allInfo.Allergies = "NONE"
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allInfo)
}

func AddAllergy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var info Info

	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON"))
		return
	}

	var userInfo Info
	result := InfoDB.First(&userInfo, "email = ?", info.Email)
	// if info is not found, create an entry for the user
	if result.Error != nil {
		InfoDB.Create(&info)
		json.NewEncoder(w).Encode(GenerateResponse("Allergy successfully added"))
		return
	}

	// check if user already has allergy logged, if not, add it
	allergies := userInfo.Allergies
	allergyList := strings.Split(allergies, ",")
	for i := 0; i < len(allergyList); i++ {
		if info.Allergies == allergyList[i] {
			json.NewEncoder(w).Encode(GenerateResponse("Allergy already exists"))
			return
		}
	}
	userInfo.Allergies += "," + info.Allergies
	InfoDB.Save(&userInfo)
	json.NewEncoder(w).Encode(GenerateResponse("Allergy successfully added"))
}
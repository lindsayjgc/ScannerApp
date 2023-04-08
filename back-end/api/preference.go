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

var PreferenceDB *gorm.DB

type Preference struct {
	gorm.Model
	User string `json:"user"`
	LabelType string `json:"type"`
	Name string `json:"name"`
}

func InitialPreferenceMigration() {
	PreferenceDB, err = gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

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
	PreferenceDB.AutoMigrate(&Preference{})
}

func AddPreference(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var newPref Label
	err = json.NewDecoder(r.Body).Decode(&newPref)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON body"))
	}

	// Verify that the preference is valid
	var checkPref Label
	result := LabelDB.Where("label_type = ?", newPref.LabelType).Where("name = ?", newPref.Name).Find(&checkPref)

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GenerateResponse("Label type is invalid"))
	}

	addPref := Preference{
		User: claims.Email,
		LabelType: newPref.LabelType,
		Name: newPref.Name,
	}
	fmt.Println(addPref)
	PreferenceDB.Create(&addPref)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(GenerateResponse("Preference successfully added"))
}
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

func GetPreferences(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	// Get preferences from DB by user email
	var preferences []Preference
	result := PreferenceDB.Table("preferences").Where("user = ?", claims.Email).Find(&preferences)

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(GenerateResponse("No preferences found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(preferences)
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
		return
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
	PreferenceDB.Create(&addPref)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(GenerateResponse("Preference successfully added"))
}

func DeletePreference(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	// Decode the preference to be deleted from the frontend
	var deleteLabel Label
	err = json.NewDecoder(r.Body).Decode(&deleteLabel)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON body"))
		return
	}

	var deletedPreference Preference
	result := PreferenceDB.Where("user = ?", claims.Email).Where("label_type = ?", deleteLabel.LabelType).Where("name = ?", deleteLabel.Name).Delete(&deletedPreference)

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GenerateResponse("Preference not found"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenerateResponse("Preference successfully deleted"))
}
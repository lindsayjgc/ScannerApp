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

var AllergyDB *gorm.DB

type Allergy struct {
	gorm.Model
	Email   string `json:"email"`
	Allergy string `json:"allergy"`
}

type RawAllergies struct {
	Allergies string `json:"allergies"`
}

func InitialAllergyMigration() {
	AllergyDB, err = gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

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
	AllergyDB.AutoMigrate(&Allergy{})
}

func AddAllergy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	// Decode allergies to be added and split into slice
	var rawNewAllergies RawAllergies
	err = json.NewDecoder(r.Body).Decode(&rawNewAllergies)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON"))
		return
	}
	newAllergies := strings.Split(string(rawNewAllergies.Allergies), ",")

	// Retrieve existing allergies as a slice
	var existingAllergiesSlice []string
	AllergyDB.Model(Allergy{}).Where("email = ?", claims.Email).Select("allergy").Find(&existingAllergiesSlice)

	// Convert the slice of allergies into a map for efficiency later
	existingAllergies := make(map[string]bool)
	for _, v := range existingAllergiesSlice {
		existingAllergies[string(v)] = true
	}

	// Check for existing allergies and add them to DB if not
	var addedAllergies []string
	var notAddedAllergies []string
	for _, v := range newAllergies {
		// If this new allergy does not already exist
		if !existingAllergies[v] {
			allergy := Allergy{Email: claims.Email, Allergy: v}
			addedAllergies = append(addedAllergies, v)
			AllergyDB.Create(&allergy)
		} else {
			notAddedAllergies = append(notAddedAllergies, v)
		}
	}

	res := make(map[string]string)
	res["addedAllergies"] = strings.Join(addedAllergies, ",")
	res["existingAllergies"] = strings.Join(notAddedAllergies, ",")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func DeleteAllergy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var rawAllergiesToDelete RawAllergies
	err = json.NewDecoder(r.Body).Decode(&rawAllergiesToDelete)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON"))
		return
	}
	allergiesToDelete := strings.Split(string(rawAllergiesToDelete.Allergies), ",")

	// Retrieve existing allergies as a slice
	var existingAllergiesSlice []string
	AllergyDB.Model(Allergy{}).Where("email = ?", claims.Email).Select("allergy").Find(&existingAllergiesSlice)

	// Convert the slice of allergies into a map for efficiency later
	existingAllergies := make(map[string]bool)
	for _, v := range existingAllergiesSlice {
		existingAllergies[string(v)] = true
	}

	var deletedAllergies []string
	var notDeletedAllergies []string
	for _, v := range allergiesToDelete {
		// If this allergy does exist, delete it
		if existingAllergies[v] {
			deletedAllergies = append(deletedAllergies, v)
			allergy := Allergy{Email: claims.Email, Allergy: v}
			AllergyDB.Where("email = ? AND allergy = ?", allergy.Email, allergy.Allergy).Delete(&allergy)
		} else {
			notDeletedAllergies = append(notDeletedAllergies, v)
		}
	}

	res := make(map[string]string)
	res["deletedAllergies"] = strings.Join(deletedAllergies, ",")
	res["notDeletedAllergies"] = strings.Join(notDeletedAllergies, ",")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

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

// Used for storing and retrieving allergies in DB with GORM
type Allergy struct {
	gorm.Model
	Email string `json:"email"`
	Allergy string `json:"allergy"`
}

// Following structs are for decoding JSON from frontend
type RawNewAllergies struct {
	Allergies string `json:"allergies"`
}

type RawProductIngredients struct {
	Ingredients string `json:"ingredients"`
}

type AllUserInfo struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	// Password  string `json:"password"`
	Allergies string `json:"allergies"`
}

type Email struct {
	Email string `json:"email"`
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

func UserInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var user User
	result := UserDB.First(&user, "email = ?", claims.Email)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(GenerateResponse("User not found"))
		return
	}

	// Retrieve user allergies as a slice 
	var userAllergiesSlice []string
	result = AllergyDB.Model(Allergy{}).Where("email = ?", claims.Email).Select("allergy").Find(&userAllergiesSlice)

	// all important user info combined into one struct for easier use by frontend
	var allInfo AllUserInfo
	allInfo.FirstName = user.FirstName
	allInfo.LastName = user.LastName
	allInfo.Email = user.Email
	// allInfo.Password = user.Password
	if len(userAllergiesSlice) == 0 {
		allInfo.Allergies = "NONE"
	} else {
		allInfo.Allergies =  strings.Join(userAllergiesSlice, ",")
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allInfo)
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
	var rawNewAllergies RawNewAllergies
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
		if !existingAllergies[v] { // If this new allergy does not already exist
			allergy := Allergy{Email: claims.Email, Allergy: v}
			addedAllergies = append(addedAllergies, v)
			AllergyDB.Create(&allergy)
		} else { // Otherwise, add to a list of already existing allergies
			notAddedAllergies = append(notAddedAllergies, v)
		}
	}

	res := make(map[string]string)
	res["addedAllergies"] = strings.Join(addedAllergies, ",")
	res["existingAllergies"] = strings.Join(notAddedAllergies, ",")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func CheckAllergies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	// Retrieve user's allergies as a slice 
	// type allergy string
	var userAllergiesSlice []string
	result := AllergyDB.Model(Allergy{}).Where("email = ?", claims.Email).Select("allergy").Find(&userAllergiesSlice)

	// Handle possible errors, this does not include if no allergies were found
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error searching for user allergies"))
	}

	// Convert the slice of allergies into a map for efficiency later
	userAllergies := make(map[string]bool)
	for _, v := range userAllergiesSlice {
		userAllergies[string(v)] = true
	}

	// Get the product allergies sent by frontend
	var rawProductIngredients RawProductIngredients
	err = json.NewDecoder(r.Body).Decode(&rawProductIngredients)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON body"))
		return
	}
	productIngredients := strings.Split(rawProductIngredients.Ingredients, ",")

	
	var foundAllergies []string
	// Compare product ingredients to user allergies
	for _, v := range productIngredients {
		if (userAllergies[v]) {
			foundAllergies = append(foundAllergies, v)
		}
	}

	if (len(foundAllergies) == 0) {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(GenerateResponse("No allergies present in product ingredients"))
	} else {
		w.WriteHeader(http.StatusOK)
		res := make(map[string]string)
		res["allergiesPresent"] = "true"
		res["allergies"] = strings.Join(foundAllergies, ",")
		json.NewEncoder(w).Encode(res)
	}
}
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

type RawNewAllergies struct {
	Allergies string `json:"allergies"`
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
	result = AllergyDB.Model(Allergy{}).Where("email = ?", claims.Email).Select("allergies").Find(&userAllergiesSlice)

	// all important user info combined into one struct for easier use by frontend
	var allInfo AllUserInfo
	allInfo.FirstName = user.FirstName
	allInfo.LastName = user.LastName
	allInfo.Email = user.Email
	// allInfo.Password = user.Password
	if len(userAllergiesSlice) == 0 {
		allInfo.Allergies = "NONE"
	} else {
		allInfo.Allergies = strings.Join(userAllergiesSlice, ",")
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
	// var info Info

	// err := json.NewDecoder(r.Body).Decode(&info)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	res := make(map[string]string)
	// 	res["msg"] = "Cannot decode user info"
	// 	json.NewEncoder(w).Encode(res)
	// 	return
	// }

	// var userInfo Info
	// result := DB.First(&userInfo, "email = ?", info.Email)
	// // if current user has no info, just ignore request
	// if result.Error != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	res := make(map[string]string)
	// 	res["msg"] = "Allergy not present for deletion"
	// 	json.NewEncoder(w).Encode(res)
	// 	return
	// }

	// allergies := userInfo.Allergies
	// allergyList := strings.Split(allergies, ",")
	// for i := 0; i < len(allergyList); i++ {
	// 	if info.Allergies == allergyList[i] {
	// 		allergyList = append(allergyList[:i], allergyList[i+1:]...)
	// 		allergies = strings.Join(allergyList, ",")
	// 		userInfo.Allergies = allergies
	// 		DB.Save(&userInfo)
	// 		res := make(map[string]string)
	// 		res["msg"] = "Allergy successfully deleted"
	// 		json.NewEncoder(w).Encode(res)
	// 		return
	// 	}
	// }

	// w.WriteHeader(http.StatusInternalServerError)
	// res := make(map[string]string)
	// res["msg"] = "Allergy not present for deletion"
	// json.NewEncoder(w).Encode(res)

}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var FavoriteDB *gorm.DB

// Used for storing and retrieving allergies in DB with GORM
type Favorite struct {
	gorm.Model
	Email   string `json:"email"`
	Favorite string `json:"favorite"`
	Code string `json:"code"`
	Image string `json:"image"`
}

// Used for receiving a new favorite from frontend
type RawFavorite struct {
	Favorite string `json:"favorite"`
	Code string `json:"code"`
	Image string `json:"image"`
}

// Used for receiving a product code and checking
// whether the product is already a favorite
// or to delete the product
type ProductCode struct {
	Code string `json:"code"`
}

func InitialFavoriteMigration() {
	FavoriteDB, err = gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

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
	FavoriteDB.AutoMigrate(&Favorite{})
}

func GetFavorites(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var favorites []RawFavorite
	result := FavoriteDB.Table("favorites").Where("email = ?", claims.Email).Where("deleted_at IS NULL").Select("favorite, code, image").Find(&favorites)

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(GenerateResponse("No favorites found"))
		return
	}

	jsonFavorites, err := json.Marshal(favorites)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error encoding JSON response"))
	}	

	w.WriteHeader(http.StatusOK)
	w.Write(jsonFavorites)
}

func CheckFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var code ProductCode
	err = json.NewDecoder(r.Body).Decode(&code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON body"))
		return
	}

	// Search for the code in the database
	var codeSearch Favorite
	result := FavoriteDB.Where("code = ?", code.Code).Where("email = ?", claims.Email).First(&codeSearch)

	// Create and assign a string for JSON response
	var isFavorite string
	if result.RowsAffected == 0 {
		isFavorite = "false"
	} else {
		isFavorite = "true"
	}

	w.WriteHeader(http.StatusOK)
	res := make(map[string]string)
	res["code"] = code.Code
	res["isFavorite"] = isFavorite
	json.NewEncoder(w).Encode(res)
}

func AddFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var newFavorite RawFavorite
	err = json.NewDecoder(r.Body).Decode(&newFavorite)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON body"))
		return
	}

	// Check if product already exists
	var checkFavorite Favorite
	result := FavoriteDB.Where("email = ?", claims.Email).Where("code = ?", newFavorite.Code).First(&checkFavorite)

	if result.RowsAffected != 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GenerateResponse("Product is already favorited"))
		return
	}

	favorite := Favorite{
		Email: claims.Email,
		Favorite: newFavorite.Favorite,
		Code: newFavorite.Code,
		Image: newFavorite.Image,
	}
	FavoriteDB.Create(&favorite)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(GenerateResponse("Product successfully favorited"))
}

func DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var code ProductCode
	err = json.NewDecoder(r.Body).Decode(&code)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON body"))
		return
	}

	var deletedFavorite Favorite
	result := FavoriteDB.Where("code = ?", code.Code).Where("email = ?", claims.Email).Delete(&deletedFavorite)

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GenerateResponse("Product not found in user favorites"))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenerateResponse("Favorite successfully deleted"))
}

func CheckCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var rawCode RawCode
	err = json.NewDecoder(r.Body).Decode(&rawCode)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON body"))
		return
	}

	// Fetch the correct code from the DB
	var codeSearch Code
	result := CodeDB.Where("email = ?", rawCode.Email).First(&codeSearch)
	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GenerateResponse("Email has not been issued a verification code"))
		return
	}

	// Create map for JSON response
	res := make(map[string]string)

	// Compare codes
	if codeSearch.Code != rawCode.Code {
		w.WriteHeader(http.StatusOK)
		res["isVerified"] = "false"
		res["message"] = "Incorrect code"
		json.NewEncoder(w).Encode(res)
		return
	}

	// Check if code has expired
	if codeSearch.Expiration.Before(time.Now()) {
		w.WriteHeader(http.StatusOK)
		res["isVerified"] = "false"
		res["message"] = "Verification code has expired"
		json.NewEncoder(w).Encode(res)
		return
	}

	var deletedCode Code
	CodeDB.Where("email = ?", rawCode.Email).Delete(&deletedCode)
	w.WriteHeader(http.StatusOK)
	res["isVerified"] = "true"
	res["message"] = "Email successfully verified"
	json.NewEncoder(w).Encode(res)
}
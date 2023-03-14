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

var FavoriteDB *gorm.DB

// Used for storing and retrieving allergies in DB with GORM
type Favorite struct {
	gorm.Model
	Email   string `json:"email"`
	Favorite string `json:"favorite"`
	Image string `json:"image"`
}

type NewFavorite struct {
	Favorite string `json:"favorite"`
	Image string `json:"image"`
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

func AddFavorite(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var newFavorite NewFavorite
	err = json.NewDecoder(r.Body).Decode(&newFavorite)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON body"))
		return
	}

	// Check if product already exists
	var checkFavorite Favorite
	result := FavoriteDB.Where("email = ?", claims.Email).Where("favorite = ?", newFavorite.Favorite).First(&checkFavorite)

	if result.RowsAffected != 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GenerateResponse("Product is already favorited"))
		return
	}

	favorite := Favorite{
		Email: claims.Email,
		Favorite: newFavorite.Favorite,
		Image: newFavorite.Image,
	}
	FavoriteDB.Create(&favorite)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(GenerateResponse("Product successfully favorited"))
}
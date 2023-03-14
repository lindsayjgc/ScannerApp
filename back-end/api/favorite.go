package main

import (
	"fmt"
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

func InitialFavoriteMigration() {
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
	FavoriteDB.AutoMigrate(&Allergy{})
}
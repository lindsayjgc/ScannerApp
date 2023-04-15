package main

import (
	"fmt"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Search struct {
	gorm.Model
	Email string `json:"email"`
	Query string `json:"query"`
	Count int    `json:"count"`
}

var SearchDB *gorm.DB

// saves searches
// if query doesn't return any products,
// check if it returns anything when searched by nutrition
// if not, don't add to database
func InitialSearchMigration() {
	SearchDB, err = gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Error connecting to Search Database.")
	}

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
		panic("Error loading .env file.")
	}

	jwtKey = []byte(os.Getenv("SECRET_KEY"))

	SearchDB.AutoMigrate(&Search{})
}

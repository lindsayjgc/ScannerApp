package main

import (
	"fmt"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Info struct {
	gorm.Model
	Email     string `json:"email"`
	Allergies string `json:"allergies"`
}

func InitialInfoMigration() {
	DB, err = gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

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
	DB.AutoMigrate(&Info{})
}

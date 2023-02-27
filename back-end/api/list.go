package main

import (
	"fmt"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var ListDB *gorm.DB

type List struct {
	gorm.Model
	Email    string `json:email`
	ListName string `json:list-name`
	Food     string `json:food`
}

func InitialListMigration() {
	ListDB, err = gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

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

	ListDB.AutoMigrate(&List{})
}

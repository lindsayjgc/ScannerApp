package main

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error
const DB_PATH = "../db/groceryapp.db"

type User struct {
	gorm.Model // Declare this as the schema for GORM
	FirstName string `json:"firstname"`
	LastName string `json:"firstname"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func InitialMigration() {
	DB, err = gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Error connecting to DB.")
	}

	// AutoMigrate checks the DB for a matching existing schema - if it does
	// not exist, create/update the new schema
	DB.AutoMigrate(&User{})
}
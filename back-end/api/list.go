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

var ListDB *gorm.DB

type List struct {
	gorm.Model
	Email string `json:"email"`
	Title string `json:"title"`
	Food  string `json:"food"`
}

type Title struct {
	gorm.Model
	Email string `json:"email"`
	Title string `json:"title"`
}

type ListTitle struct {
	Title string `json:"title"`
}

func InitialListMigration() {
	ListDB, err = gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

	AllergyDB.Exec("DROP TABLE lists")

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
	ListDB.AutoMigrate(&Title{})
}

func CreateList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var listTitle ListTitle
	err = json.NewDecoder(r.Body).Decode(&listTitle)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON"))
		return
	}

	// Check if title already exists in database
	var existingTitle Title
	res := ListDB.First(&existingTitle, "email = ? AND title = ?", claims.Email, listTitle.Title)

	if res.Error == nil {
		// title already exists in db
		w.WriteHeader(http.StatusAlreadyReported)
		json.NewEncoder(w).Encode(GenerateResponse("List already exists"))
		return
	}

	var title Title
	title.Email = claims.Email
	title.Title = listTitle.Title
	ListDB.Create(&title)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(GenerateResponse("List successfully created"))

}

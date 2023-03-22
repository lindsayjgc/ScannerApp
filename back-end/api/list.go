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

type GroceryItem struct {
	gorm.Model
	Email string `json:"email"`
	Title string `json:"title"`
	Item  string `json:"item"`
}

type GroceryTitle struct {
	gorm.Model
	Email string `json:"email"`
	Title string `json:"title"`
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

	ListDB.AutoMigrate(&GroceryItem{})
	ListDB.AutoMigrate(&GroceryTitle{})
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

	var listTitle GroceryTitle
	err = json.NewDecoder(r.Body).Decode(&listTitle)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON"))
		return
	}

	// Check if title already exists in database
	var existingTitle GroceryTitle
	res := ListDB.First(&existingTitle, "email = ? AND title = ?", claims.Email, listTitle.Title)

	if res.Error == nil {
		// title already exists in db
		w.WriteHeader(http.StatusAlreadyReported)
		json.NewEncoder(w).Encode(GenerateResponse("List already exists"))
		return
	}

	listTitle.Email = claims.Email
	ListDB.Create(&listTitle)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(GenerateResponse("List successfully created"))

}

func AddGroceryItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var listItem GroceryItem
	err = json.NewDecoder(r.Body).Decode(&listItem)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON"))
		return
	}

	// Check if item is already on list
	var existingItem GroceryItem
	res := ListDB.First(&existingItem, "email = ? AND title = ? AND item = ?", claims.Email, listItem.Title, listItem.Item)

	if res.Error == nil {
		// item already exists in db
		w.WriteHeader(http.StatusAlreadyReported)
		json.NewEncoder(w).Encode(GenerateResponse("Item already exists"))
		return
	}

	listItem.Email = claims.Email
	ListDB.Create(&listItem)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(GenerateResponse("Item successfully added"))
}

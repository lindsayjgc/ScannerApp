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

type Search struct {
	gorm.Model
	Email string `json:"email"`
	Query string `json:"query"`
	Count int    `json:"count"`
}

type RawSearch struct {
	Query string `json:"query"`
}

type Query struct {
	Query string
	Count int
}

var SearchDB *gorm.DB

// saves searches
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

func SaveQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var query RawSearch
	err = json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON"))
		return
	}

	var existingQuery Search
	res := SearchDB.First(&existingQuery, "email = ? AND query = ?", claims.Email, query.Query)

	if res.Error == nil {
		existingQuery.Count = existingQuery.Count + 1
		SearchDB.Save(existingQuery)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(GenerateResponse("Query count updated"))
		return
	}

	// TODO: check if there is a response from the query
	search := Search{Email: claims.Email, Query: query.Query, Count: 1}
	SearchDB.Create(search)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(GenerateResponse("Query count updated"))
}

func GetQueries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var existingQueriesSlice []Query
	SearchDB.Model(&Search{}).Select("query, count").Where("email = ?", claims.Email).Group("query").Scan(&existingQueriesSlice)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingQueriesSlice)
}

func RemoveQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var deleteSearch RawSearch
	err = json.NewDecoder(r.Body).Decode(&deleteSearch)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON"))
		return
	}

	var query Search
	res := SearchDB.Where("email = ? AND query = ?", claims.Email, deleteSearch.Query).First(&query)
	if res.Error != nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(GenerateResponse("Query not deleted"))
		return
	}
	SearchDB.Where("email = ? AND query = ?", query.Email, query.Query).Delete(&query)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenerateResponse("Query successfully deleted"))
}

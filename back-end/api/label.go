package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var LabelDB *gorm.DB

type Label struct {
	gorm.Model
	LabelType string `json:"type"`
	Name      string `json:"name"`
}

func InitialLabelMigration() {
	LabelDB, err = gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Error connecting to UserDB.")
	}

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
		panic("Error loading .env file.")
	}

	jwtKey = []byte(os.Getenv("SECRET_KEY"))

	// AutoMigrate checks the DB for a matching existing schema - if it does
	// not exist, create/update the new schema
	LabelDB.AutoMigrate(&Label{})
}

func GetLabels(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get the desired label type from query params
	labelType := r.URL.Query().Get("type")
	
	if labelType == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GenerateResponse("Missing query parameter: type"))
		return
	}

	// Check that label type from query params 
	labelTypes := []string{"Diet", "Health", "mealType", "dishType"}
	isValidLabel := false
	for _, l := range labelTypes {
		if l == labelType {
			isValidLabel = true
			break
		}
	}

	if !isValidLabel {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GenerateResponse("Label type is invalid"))
		return
	}

	// Get the labels
	var labels []Label
	result := LabelDB.Table("labels").Omit("Date").Select("name").Find(&labels, "label_type = ?", labelType)
	if result.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error retrieving labels from DB. You may need to add the labels to your instance of the DB"))
		return
	}

	// Create an array of just the label names
	var labelNames []string
	for _, l := range labels {
		labelNames = append(labelNames, l.Name)
	}

	// Create a comma-delimited string of the labels
	labelString := strings.Join(labelNames, ",")

	// Encode the string in the response
	w.WriteHeader(http.StatusOK)
	res := make(map[string]string)
	res["labels"] = labelString
	json.NewEncoder(w).Encode(res)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var RecipeDB *gorm.DB

type Recipe struct {
	gorm.Model
	User string
	APIURI string
	SourceURL string
	RecipeID string
	Label string
	Liked *bool // Denotes whether a user has liked or disliked a recommendation, can be null
}

type Response struct {
	From  int `json:"from"`
	To    int `json:"to"`
	Count int `json:"count"`
	Links struct {
			Self struct {
					Href  string `json:"href"`
					Title string `json:"title"`
			} `json:"self"`
			Next struct {
					Href  string `json:"href"`
					Title string `json:"title"`
			} `json:"next"`
	} `json:"_links"`
	Hits []struct {
			Recipe struct {
					Uri              string `json:"uri"`
					Label            string `json:"label"`
					Image            string `json:"image"`
					Images           map[string]struct {
							Url    string `json:"url"`
							Width  int    `json:"width"`
							Height int    `json:"height"`
					} `json:"images"`
					Source            string    `json:"source"`
					Url               string    `json:"url"`
					ShareAs           string    `json:"shareAs"`
					// Yield             int    		`json:"yield,omitempty"`
					DietLabels        []string  `json:"dietLabels"`
					HealthLabels      []string  `json:"healthLabels"`
					Cautions          []string  `json:"cautions"`
					IngredientLines   []string  `json:"ingredientLines"`
					Ingredients       []struct {
							Text     string  `json:"text"`
							Quantity float64 `json:"quantity"`
							Measure  string  `json:"measure"`
							Food     string  `json:"food"`
							Weight   float64 `json:"weight"`
							FoodId   string  `json:"foodId"`
					} `json:"ingredients"`
					Calories           float64 `json:"calories"`
					GlycemicIndex      int     `json:"glycemicIndex"`
					TotalCO2Emissions  float64 `json:"totalCO2Emissions"`
					Co2EmissionsClass  string  `json:"co2EmissionsClass"`
					TotalWeight        float64 `json:"totalWeight"`
					CuisineType        []string `json:"cuisineType"`
					MealType           []string `json:"mealType"`
					DishType           []string `json:"dishType"`
					Instructions       []string `json:"instructions"`
					Tags               []string `json:"tags"`
					// ExternalId         string   `json:"externalId"`
					// TotalNutrients     struct{} `json:"totalNutrients"`
					// TotalDaily         struct{} `json:"totalDaily"`
					// Digest             []struct {
					// 		Label        string  `json:"label"`
					// 		Tag          string  `json:"tag"`
					// 		SchemaOrgTag string  `json:"schemaOrgTag"`
					// 		Total        float64 `json:"total"`
					// 		HasRDI       bool    `json:"hasRDI"`
					// 		Daily        float64 `json:"daily"`
					// 		Unit         string  `json:"unit"`
							// Sub          struct{} `json:"sub"`
					// } `json:"digest"`
			} `json:"recipe"`
			// Links struct {
			// 		Self struct {
			// 				Href  string `json:"href"`
			// 				Title string `json:"title"`
			// 		} `json:"self"`
			// 		Next struct {
			// 				Href  string `json:"href"`
			// 				Title string `json:"title"`
			// 		} `json:"next"`
			// } `json:"_links"`
	} `json:"hits"`
}


func InitialRecipeMigration() {
	RecipeDB, err = gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

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
	RecipeDB.AutoMigrate(&Recipe{})
}

func GetRecipeRecommendations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	// Get the distinct preference label types that the user has
	// Used for formulating the search URL for the recipe API
	// var preferenceTypes []string
	// result := PreferenceDB.Table("preferences").Distinct("label_type").Where("user = ?", claims.Email).Where("deleted_at IS NULL").Pluck("label_type", &preferenceTypes)

	// if result.RowsAffected == 0 {
	// 	w.WriteHeader(http.StatusNoContent)
	// 	json.NewEncoder(w).Encode(GenerateResponse("User has not indicated any preferences for recipes"))
	// 	return
	// }

	// Get all of the user's preferences for use in the search URL
	var preferences []Preference
	result := PreferenceDB.Table("preferences").Where("user = ?", claims.Email).Find(&preferences)

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(GenerateResponse("User has not indicated any preferences for recipes"))
		return
	}

	// Create search URL for API
	API_URL := os.Getenv("RECIPE_API_BASE_URL")
	for _, p := range preferences {
		API_URL = API_URL + "&" + p.LabelType + "=" + p.Name
	}

	// Create a new HTTP client and GET request
	client := &http.Client{}
	req, err := http.NewRequest("GET", API_URL, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error creating request to API"))
		return
	}

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error sending request to API"))
		return
	}
	defer resp.Body.Close() // Close body when done

	// Read response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error reading response from API"))
		return
	}

	// Unmarshal in an object
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error translating API response into JSON"))
		panic(err)
	}

	// Process the recipes into the DB and an array for return
	var recipeHits []Recipe
	for _, rec := range response.Hits {
		newRecipe := Recipe{
			User: claims.Email,
			APIURI: rec.Recipe.Uri,
			SourceURL: rec.Recipe.Url,
			// Capture the ID from after the # in the URI
			RecipeID: strings.Split(rec.Recipe.Uri, "#")[1], 
			Label: rec.Recipe.Label,
			Liked: nil,
		}

		recipeHits = append(recipeHits, newRecipe)
		RecipeDB.Save(&newRecipe)
	}
	
	// Respond with the recipes
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recipeHits)
}

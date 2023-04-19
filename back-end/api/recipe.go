package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	Liked string
}

// Used to decode data from frontend when trying to change 
// whether a user like/dislikes a recipe
type LikeStatusData struct {
	RecipeID string `json:"id"`
	Liked string `json:"liked"`
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
			} `json:"recipe"`
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
	
	// The handler should always return 10 recipes; use a loop structure, as we
	// process out recipes that have already been liked/disliked by the user
	// Because of this, we may need to make multiple API calls to get more recipes
	var recipeHits []Recipe // Recipes that will be returned to the user
	const RECIPE_MAX = 5

	for len(recipeHits) < RECIPE_MAX {
		req, err := http.NewRequest("GET", API_URL, nil)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(GenerateResponse("Error creating request to API"))
			return
		}

		// Send request
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(GenerateResponse("Error sending request to API"))
			return
		}
		defer resp.Body.Close() // Close body when done

		// Read response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(GenerateResponse("Error reading response from API"))
			return
		}

		// Unmarshal in an object
		var response Response
		err = json.Unmarshal(body, &response)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(GenerateResponse("Error translating API response into JSON"))
			panic(err)
		}

		// Process the recipes into the DB and an array for return
		for _, rec := range response.Hits {
			// Capture the ID from after the # in the URI
			newRecipeID := strings.Split(rec.Recipe.Uri, "#")[1]
			
			// Check if this recipe has already been liked/disliked by the user
			var recipeSearch Recipe
			result := RecipeDB.Table("recipes").Where("user = ?", claims.Email).Where("recipe_id = ?", newRecipeID).First(&recipeSearch)
			if result.RowsAffected != 0 && recipeSearch.Liked != "none" {
				continue
			}

			// Create new recipe object from API info
			newRecipe := Recipe{
				User: claims.Email,
				APIURI: rec.Recipe.Uri,
				SourceURL: rec.Recipe.Url,
				RecipeID: newRecipeID,
				Label: rec.Recipe.Label,
				Liked: "none",
			}
			recipeHits = append(recipeHits, newRecipe)

			// Only save the recipe if it wasn't found in the search
			if result.RowsAffected == 0 {
				RecipeDB.Save(&newRecipe)
			}

			if len(recipeHits) >= RECIPE_MAX {
				break
			}
		}

		// Check if we have the criteria of 5 recipes to be returned
		// If so, change the API_URL for the next request to the next
		// page of results
		if len(recipeHits) < RECIPE_MAX {
			API_URL = response.Links.Next.Href
		}
	}
	
	// Respond with the recipes
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(recipeHits)
}

func UpdateRecipeLikeStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	// Decode the recipe to update
	var updateStatus LikeStatusData
	err = json.NewDecoder(r.Body).Decode(&updateStatus)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON body"))
		return
	}

	// Check that the liked status is invalid
	if updateStatus.Liked != "true" && updateStatus.Liked != "false" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GenerateResponse("Invalid like status"))
		return
	}

	var recipeSearch Recipe
	result := RecipeDB.Where("user = ?", claims.Email).Where("recipe_id", updateStatus.RecipeID).First(&recipeSearch)

	if result.RowsAffected == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GenerateResponse("Recipe not found in user recommendations"))
		return
	}

	if recipeSearch.Liked == updateStatus.Liked {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(GenerateResponse("Like status already set to " + recipeSearch.Liked))
		return		
	}

	recipeSearch.Liked = updateStatus.Liked
	RecipeDB.Save(&recipeSearch)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(GenerateResponse("Recommendation status updated"))
}
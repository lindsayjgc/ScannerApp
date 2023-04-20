package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetRecipeRecommendations(t *testing.T) {
	InitialPreferenceMigration()
	InitialRecipeMigration()
	InitializeRouter()

	// Add a preference for recommendations to be based on
	testPref := Preference{
		User: "unit@test.com",
		LabelType: "diet",
		Name: "balanced",
	}
	PreferenceDB.Create(&testPref)

	req, _ := http.NewRequest("GET", "/api/recipe/recommendation", nil)
	req.Header.Set("Content-Type", "application/json")

	// Create test cookie to simulate user login
	createTestCookie(req, t)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Recommendations change so just check the status, 200 means recommendations were returned properly
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned the wrong status: got %v, expected %v", status, http.StatusOK)
	}

	// Clean up test data
	var deletedPref Preference
	PreferenceDB.Where("user = ?", "unit@test.com").Unscoped().Delete(&deletedPref)
	var deletedRecipes []Recipe
	RecipeDB.Table("recipes").Where("user = ?", "unit@test.com").Unscoped().Delete(&deletedRecipes)
}

func TestUpdateRecipeLikeStatus(t *testing.T) {
	InitialPreferenceMigration()
	InitialRecipeMigration()
	InitializeRouter()

	// Create a test recipe to be updated
	testRecipe := Recipe{
		User: "unit@test.com",
		APIURI: "testuri.com",
		SourceURL: "testurl.com",
		RecipeID: "testid",
		Label: "testlabel",
		Liked: "none",
	}
	RecipeDB.Create(&testRecipe)

	testLikeStatus := LikeStatusData{
		RecipeID: "testid",
		Liked: "true",
	}

	payload, _ := json.Marshal(testLikeStatus)
	req, _ := http.NewRequest("PUT", "/api/recipe/update", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// Create test cookie to simulate user login
	createTestCookie(req, t)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Recommendations change so just check the status, 200 means recommendations were returned properly
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned the wrong status: got %v, expected %v", status, http.StatusOK)
	}

	expected := `{"message":"Recommendation status updated"}`
	// Remove any linebreaks from the response writer body
	body := strings.Replace(rr.Body.String(), "\n", "", -1)
	body = strings.Replace(body, "\r", "", -1)

	if body != expected {
		t.Errorf("Handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected)
	}

	// Clean up test data
	var deletedRecipes []Recipe
	RecipeDB.Table("recipes").Where("user = ?", "unit@test.com").Unscoped().Delete(&deletedRecipes)
}
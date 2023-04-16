package main

import (
	"net/http"
	"net/http/httptest"
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
package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestGetPreferences(t *testing.T) {
	InitialPreferenceMigration()
	InitialLabelMigration()
	InitializeRouter()

	testPref := Preference{
		User: "unit@test.com",
		LabelType: "diet",
		Name: "balanced",
	}
	PreferenceDB.Create(&testPref)

	// Create a mock request
	req, _ := http.NewRequest("GET", "/api/preference", nil)
	req.Header.Set("Content-Type", "application/json")
	
	// Create test cookie to simulate user login
	createTestCookie(req, t)

	// Create a mock request recorder
	rr := httptest.NewRecorder()

	// Send the request and recorder
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned the wrong status: got %v, expected %v", status, http.StatusCreated)
	}

	expectedRegex := regexp.MustCompile(`(?i)"user"\s*:\s*"unit@test.com"\s*,\s*"type"\s*:\s*"diet"\s*,\s*"name"\s*:\s*"balanced"`)
	// body := strings.Replace(rr.Body.String(), "\n", "", -1);
	// body = strings.Replace(body, "\r", "", -1);

	if !expectedRegex.MatchString(rr.Body.String()) {
		t.Errorf("Handler returned unexpected body: got %v", rr.Body.String());
	}

	// Clean up test data from DB
	var deletedPrefs []Preference
	FavoriteDB.Table("preferences").Where("email = ?", "unit@test.com").Unscoped().Delete(&deletedPrefs)
	fmt.Println(deletedPrefs)
}
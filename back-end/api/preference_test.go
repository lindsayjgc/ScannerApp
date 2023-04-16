package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
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
	PreferenceDB.Table("preferences").Where("user = ?", "unit@test.com").Unscoped().Delete(&deletedPrefs)
}

func TestAddPreference(t *testing.T) {
	InitialPreferenceMigration()
	InitialLabelMigration()
	InitializeRouter()

	newPref := Label{
		LabelType: "diet",
		Name: "balanced",
	}

	payload, _ := json.Marshal(newPref)
	req, _ := http.NewRequest("POST", "/api/preference", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// Create test cookie to simulate user login
	createTestCookie(req, t)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned the wrong status: got %v, expected %v", status, http.StatusCreated)
	}

	expected := `{"message":"Preference successfully added"}`
	body := strings.Replace(rr.Body.String(), "\n", "", -1)
	body = strings.Replace(body, "\r", "", -1)

	if body != expected {
		t.Errorf("Handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected)
	}

	// Delete the user that was created for the test
	var deletedPref Preference
	PreferenceDB.Where("user = ?", "unit@test.com").Unscoped().Delete(&deletedPref)
}
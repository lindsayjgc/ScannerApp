package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCheckAllerigesNotFound(t *testing.T) {
	// Initialize router and connect to DB for this test instance
	InitialUserMigration()
	InitialAllergyMigration()
	InitializeRouter()

	ingredients := RawProductIngredients {
		Ingredients: "walnuts",
	}

	payload, _ := json.Marshal(ingredients);
	req, _ := http.NewRequest("POST", "/api/check-allergies", bytes.NewBuffer(payload));
	req.Header.Set("Content-Type", "application/json")
	createTestCookie(req, t)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr,req)

	// Process response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned the wrong status: got %v, expected %v", status, http.StatusCreated);
	}

	expected := `{"allergiesPresent":"false"}`
	body := strings.Replace(rr.Body.String(), "\n", "", -1);
	body = strings.Replace(body, "\r", "", -1);

	if body != expected {
		t.Errorf("Handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected);
	}
}

// User "unit@test.com" has allergies milk, eggs, and soy stored in the database
func TestCheckAllerigesFound(t *testing.T) {
	// Initialize router and connect to DB for this test instance
	InitialUserMigration()
	InitialAllergyMigration()
	InitializeRouter()

	ingredients := RawProductIngredients {
		Ingredients: "milk,eggs,soy",
	}

	payload, _ := json.Marshal(ingredients);
	req, _ := http.NewRequest("POST", "/api/check-allergies", bytes.NewBuffer(payload));
	req.Header.Set("Content-Type", "application/json")
	createTestCookie(req, t)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr,req)

	// Process response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned the wrong status: got %v, expected %v", status, http.StatusCreated);
	}

	expected := `{"allergies":"milk,eggs,soy","allergiesPresent":"true"}`
	body := strings.Replace(rr.Body.String(), "\n", "", -1);
	body = strings.Replace(body, "\r", "", -1);

	if body != expected {
		t.Errorf("Handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected);
	}
}
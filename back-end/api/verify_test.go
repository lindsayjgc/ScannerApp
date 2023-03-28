package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestVerifyEmailSignup(t *testing.T) {
	// Initialize router and connect to DB for this test instance
	InitialUserMigration()
	InitialCodeMigration()
	InitializeRouter()

	// Specify email for verification
	email := Email{
		Email: "cen3031groceryapp@gmail.com",
	}

	// Create a mock request
	payload, _ := json.Marshal(email)
	req, _ := http.NewRequest("POST", "/api/verify/signup", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// Create a mock request recorder
	rr := httptest.NewRecorder()

	// Send the request and recorder
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned the wrong status: got %v, expected %v", status, http.StatusCreated)
	}

	expected := `{"message":"Verification email sent successfully"}`

	// Remove any linebreaks from the response writer body
	body := strings.Replace(rr.Body.String(), "\n", "", -1)
	body = strings.Replace(body, "\r", "", -1)

	if body != expected {
		t.Errorf("Handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected)
	}

	// Delete the code that was created by the API request
	var deletedCode Code
	CodeDB.Where("email = ?", "cen3031groceryapp@gmail.com").Unscoped().Delete(&deletedCode)
}
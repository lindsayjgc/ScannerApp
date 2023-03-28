package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
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

func TestVerifyReset(t *testing.T) {
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
	req, _ := http.NewRequest("POST", "/api/verify/reset", bytes.NewBuffer(payload))
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

func TestCheckCode(t *testing.T) {
	// Initialize router and connect to DB for this test instance
	InitialUserMigration()
	InitialCodeMigration()
	InitializeRouter()

	// Add code to DB To be checked
	code := Code{
		Email: "unit@test.com",
		Code: "000000",
		Expiration: time.Now().Add(time.Minute),
	}
	CodeDB.Create(&code)

	rawCode := RawCode{
		Email: "unit@test.com",
		Code: "000000",
	}

	payload, _ := json.Marshal(rawCode)
	req, _ := http.NewRequest("POST", "/api/check-code", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// Create a mock request recorder
	rr := httptest.NewRecorder()

	// Send the request and recorder
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned the wrong status: got %v, expected %v", status, http.StatusCreated)
	}

	expected := `{"isVerified":"true","message":"Email successfully verified"}`

	// Remove any linebreaks from the response writer body
	body := strings.Replace(rr.Body.String(), "\n", "", -1)
	body = strings.Replace(body, "\r", "", -1)

	if body != expected {
		t.Errorf("Handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected)
	}

	// Delete the user that was created for the test
	var deletedCode Code
	CodeDB.Where("email = ?", "unit@test.com").Unscoped().Delete(&deletedCode)
}
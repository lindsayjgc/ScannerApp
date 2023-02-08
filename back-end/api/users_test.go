package main

import (
	"bytes"
	"encoding/json"
	// "fmt"
	"net/http"
	"net/http/httptest"
	// "reflect"
	"strings"
	"testing"
)

func TestSignUpEndpoint(t *testing.T) {
	// Initialize router and connect to DB for this test instance
	InitialUserMigration()
	InitialInfoMigration()
	InitializeRouter()

	// Create a user to be added
	user := User{
		FirstName: "testfirstname",
		LastName: "testlastname",
		Email: "unit@test.com",
		Password: "unittest",
	}

	// Create a mock request 
	payload, _ := json.Marshal(user);
	req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload));
	req.Header.Set("Content-Type", "application/json")

	// Create a mock request recorder
	rr := httptest.NewRecorder()

	// Send the request and recorder
	r.ServeHTTP(rr,req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned the wrong status: got %v, expected %v", status, http.StatusCreated);
	}

	expected := `{"message":"User successfully created"}`

	// Remove any linebreaks from the response writer body
	body := strings.Replace(rr.Body.String(), "\n", "", -1);
	body = strings.Replace(body, "\r", "", -1);

	if body != expected {
		t.Errorf("Handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected);
	}
}
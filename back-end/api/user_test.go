package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

func TestSignUpEndpoint(t *testing.T) {
	// Initialize router and connect to DB for this test instance
	InitialUserMigration()
	InitialAllergyMigration()
	InitializeRouter()

	// Create a user to be added
	user := User{
		FirstName: "testfirstname",
		LastName: "testlastname",
		Email: "unit@test.com",
		Password: "ut",
	}

	// Create a mock request
	payload, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/api/signup", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	// Create a mock request recorder
	rr := httptest.NewRecorder()

	// Send the request and recorder
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned the wrong status: got %v, expected %v", status, http.StatusCreated)
	}

	expected := `{"message":"User successfully created"}`

	// Remove any linebreaks from the response writer body
	body := strings.Replace(rr.Body.String(), "\n", "", -1)
	body = strings.Replace(body, "\r", "", -1)

	if body != expected {
		t.Errorf("Handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected)
	}
}

func TestLoginEndpoint(t *testing.T) {
	InitialUserMigration()
	InitialAllergyMigration()
	InitializeRouter()

	creds := Credentials{
		Email: "unit@test.com",
		Password: "ut",
	}

	payload, _ := json.Marshal(creds)
	req, _ := http.NewRequest("POST", "/api/login", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("Handler returned the wrong status: got %v, expected %v", status, http.StatusCreated)
	}

	expected := `{"message":"User successfully logged in"}`
	body := strings.Replace(rr.Body.String(), "\n", "", -1)
	body = strings.Replace(body, "\r", "", -1)

	if body != expected {
		t.Errorf("Handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected)
	}
}

func TestLoggedInEndpoint(t *testing.T) {
	InitialUserMigration()
	InitialAllergyMigration()
	InitializeRouter()

	// Make a mock request
	req, _ := http.NewRequest("GET", "/api/logged-in", nil)
	req.Header.Set("Content-Type", "application/json")

	// Add a new cookie to the req for testing the endpoint
	createTestCookie(req, t)

	// Create a new recorder and serve
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Process response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned the wrong status: got %v, expected %v", status, http.StatusCreated)
	}

	expected := `{"email":"unit@test.com","message":"User is currently logged in"}`
	body := strings.Replace(rr.Body.String(), "\n", "", -1)
	body = strings.Replace(body, "\r", "", -1)

	if body != expected {
		t.Errorf("Handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected)
	}
}

func TestUserInfo(t *testing.T) {
	TestSignUpEndpoint(t)

	req, _ := http.NewRequest("GET", "/api/user-info", nil)
	resp := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")

	createTestCookie(req, t)

	handler := http.HandlerFunc(UserInfo)
	handler.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("Error: response code should be %v, got %v", http.StatusOK, resp.Code)
	}
	// // Decode response to our dummy user
	var decodedUser AllUserInfo
	err := json.NewDecoder(resp.Body).Decode(&decodedUser)

	userInfo := AllUserInfo{
		FirstName: "testfirstname",
		LastName:  "testlastname",
		Email:     "unit@test.com",
		Allergies: "NONE",
	}

	if err != nil {
		t.Fatalf("Error: failed decoding response body")
	}

	if decodedUser.FirstName != userInfo.FirstName ||
		decodedUser.LastName != userInfo.LastName ||
		decodedUser.Email != userInfo.Email ||
		decodedUser.Allergies != "NONE" {
		t.Errorf("Error: expected %v, got %v", userInfo, decodedUser)
	}

	TestDeleteEndpoint(t)
}

func TestLogoutEndpoint(t *testing.T) {
	InitialUserMigration()
	InitialAllergyMigration()
	InitializeRouter()

	req, _ := http.NewRequest("POST", "/api/logout", nil)
	req.Header.Set("Content-Type", "application/json")
	createTestCookie(req, t)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	// Process response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned the wrong status: got %v, expected %v", status, http.StatusCreated)
	}

	expected := `{"email":"unit@test.com","message":"User successfully logged out"}`
	body := strings.Replace(rr.Body.String(), "\n", "", -1)
	body = strings.Replace(body, "\r", "", -1)

	if body != expected {
		t.Errorf("Handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected)
	}
}

func TestDeleteEndpoint(t *testing.T) {
	InitialUserMigration()
	InitialAllergyMigration()
	InitializeRouter()

	req, _ := http.NewRequest("DELETE", "/api/delete-user", nil);
	req.Header.Set("Content-Type", "application/json")
	createTestCookie(req, t)

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr,req)

	// Process response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned the wrong status: got %v, expected %v", status, http.StatusCreated);
	}

	expected := `{"email":"unit@test.com","message":"User successfully deleted"}`
	body := strings.Replace(rr.Body.String(), "\n", "", -1);
	body = strings.Replace(body, "\r", "", -1);

	if body != expected {
		t.Errorf("Handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected);
	}
}

func createTestCookie(req *http.Request, t *testing.T) {
	// Create a new token string
	expirationTime := time.Now().Add(time.Minute)
	claims := &Claims{
		Email: "unit@test.com",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		t.Fatal(err)
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Path:    "/",
		Expires: expirationTime,
	}
	req.AddCookie(cookie)
}

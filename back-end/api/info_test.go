package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserInfo(t *testing.T) {
	// Initialize router and connect to DB for this test instance
	InitialUserMigration()
	InitialInfoMigration()
	InitializeRouter()

	email := Email{
		Email: "fl@gmail.com",
	}

	payload, _ := json.Marshal(email)
	req, err := http.NewRequest("GET", "/userinfo", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder
	rr := httptest.NewRecorder()
	var allInfo AllUserInfo

	// Invoke UserInfo with the given request and response recorder
	UserInfo(rr, req)

	// Check the status code is what we expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode json
	err = json.NewDecoder(rr.Body).Decode(&allInfo)
	if err != nil {
		t.Errorf("Can't decode json response")
	}

	// Assertions about the response
	if allInfo.FirstName != "First" {
		t.Error("First name not as expected")
	}
	if allInfo.LastName != "Last" {
		t.Error("Last name not as expected")
	}
	if allInfo.Email != "fl@gmail.com" {
		t.Error("Email not as expected")
	}
	if allInfo.Password != "fl" {
		t.Error("Password not as expected")
	}
	if allInfo.Allergies != "NONE" {
		t.Error("Allergies not as expected")
	}
}

func TestAddAllergy(t *testing.T) {
	// Initialize router and connect to DB for this test instance
	InitialUserMigration()
	InitialInfoMigration()
	InitializeRouter()

	info := Info{
		Email:     "js@gmail.com",
		Allergies: "testAllergy",
	}

	payload, _ := json.Marshal(info)
	req, err := http.NewRequest("PUT", "/update-allergies", bytes.NewBuffer(payload))

	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		t.Fatal(err)
	}

	// create ResponseRecorder
	rr := httptest.NewRecorder()

	// invoke the target function
	AddAllergy(rr, req)

	// get the response
	response := rr.Result()
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	// assertion
	if http.StatusOK != response.StatusCode {
		t.Errorf("Expected %v; got %v\n", http.StatusOK, response.StatusCode)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Errorf("Error unmarshalling response: %s \n", err)
	}

	expectedErrMsg := "Allergy successfully added"
	if resp["msg"] != expectedErrMsg {
		t.Errorf("Expected %s; got %s \n", expectedErrMsg, resp["msg"])
	}

	// TODO: clean up testAllergy
}

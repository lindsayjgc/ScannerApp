package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddAllergy(t *testing.T) {
	InitialUserMigration()
	InitialAllergyMigration()
	InitializeRouter()

	testAllergy := RawAllergies{
		Allergies: "testallergy",
	}

	payload, _ := json.Marshal(testAllergy)
	req, _ := http.NewRequest("PUT", "/add-allergies", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	createTestCookie(req, t)

	// create ResponseRecorder
	rr := httptest.NewRecorder()

	// invoke the target function
	AddAllergy(rr, req)

	// get the response
	response := rr.Result()
	body, _ := ioutil.ReadAll(response.Body)

	// assertion
	if http.StatusOK != response.StatusCode {
		t.Errorf("Expected %v; got %v\n", http.StatusOK, response.StatusCode)
	}

	resMap := make(map[string]string)
	json.Unmarshal(body, &resMap)

	expectedMsg := "testallergy"
	if resMap["addedAllergies"] != expectedMsg || resMap["existingAllergies"] != "" {
		t.Errorf("Expected %s; got %s \n", expectedMsg, resMap["addedAllergies"])
	}

	// clean up testAllergy
	allergy := Allergy{Email: "unit@test.com", Allergy: "testallergy"}
	AllergyDB.Where("email = ? AND allergy = ?", allergy.Email, allergy.Allergy).Unscoped().Delete(&allergy)
}

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

// User "unit@test.com" has allergy "testallergy" stored in the database
func TestCheckAllerigesFound(t *testing.T) {
	// Initialize router and connect to DB for this test instance
	InitialUserMigration()
	InitialAllergyMigration()
	InitializeRouter()

	// Add allergy to DB temporarily
	allergy := Allergy{Email: "unit@test.com", Allergy: "testallergy"}
	AllergyDB.Create(&allergy)

	ingredients := RawProductIngredients {
		Ingredients: "testallergy",
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

	expected := `{"allergies":"testallergy","allergiesPresent":"true"}`
	body := strings.Replace(rr.Body.String(), "\n", "", -1);
	body = strings.Replace(body, "\r", "", -1);

	if body != expected {
		t.Errorf("Handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected);
	}

	// Delete dummy allergy from DB
	allergy = Allergy{Email: "unit@test.com", Allergy: "testallergy"}
	AllergyDB.Where("email = ? AND allergy = ?", allergy.Email, allergy.Allergy).Unscoped().Delete(&allergy)
}


func TestDeleteAllergy(t *testing.T) {
	InitialUserMigration()
	InitialAllergyMigration()
	InitializeRouter()

	// create allergy to delete
	allergy := Allergy{Email: "unit@test.com", Allergy: "testallergy"}
	AllergyDB.Create(&allergy)

	testAllergy := RawAllergies{
		Allergies: "testallergy",
	}

	payload, _ := json.Marshal(testAllergy)
	req, _ := http.NewRequest("DELETE", "/delete-allergies", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	createTestCookie(req, t)

	// create ResponseRecorder
	rr := httptest.NewRecorder()

	// invoke the target function
	DeleteAllergy(rr, req)

	// get the response
	response := rr.Result()
	body, _ := ioutil.ReadAll(response.Body)

	// assertion
	if http.StatusOK != response.StatusCode {
		t.Errorf("Expected %v; got %v\n", http.StatusOK, response.StatusCode)
	}

	resMap := make(map[string]string)
	json.Unmarshal(body, &resMap)

	expectedMsg := "testallergy"
	if resMap["deletedAllergies"] != expectedMsg || resMap["notDeletedAllergies"] != "" {
		t.Errorf("Expected %s; got %s \n", expectedMsg, resMap["deletedAllergies"])
	}
}
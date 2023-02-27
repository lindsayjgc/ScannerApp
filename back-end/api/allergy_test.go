package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddAllergy(t *testing.T) {
	InitialUserMigration()
	InitialAllergyMigration()
	InitializeRouter()

	testAllergy := RawAllergies{
		Allergies: "testAllergy",
	}

	payload, _ := json.Marshal(testAllergy)
	req, _ := http.NewRequest("PUT", "/add-allergies", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	createCookie(req, t)

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

	expectedMsg := "testAllergy"
	if resMap["addedAllergies"] != expectedMsg || resMap["existingAllergies"] != "" {
		t.Errorf("Expected %s; got %s \n", expectedMsg, resMap["addedAllergies"])
	}

	// clean up testAllergy
	allergy := Allergy{Email: "unit@test.com", Allergy: "testAllergy"}
	AllergyDB.Where("email = ? AND allergy = ?", allergy.Email, allergy.Allergy).Delete(&allergy)
}

func TestDeleteAllergy(t *testing.T) {
	InitialUserMigration()
	InitialAllergyMigration()
	InitializeRouter()

	// create allergy to delete
	allergy := Allergy{Email: "unit@test.com", Allergy: "testAllergy"}
	AllergyDB.Create(&allergy)

	testAllergy := RawAllergies{
		Allergies: "testAllergy",
	}

	payload, _ := json.Marshal(testAllergy)
	req, _ := http.NewRequest("DELETE", "/delete-allergies", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	createCookie(req, t)

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

	expectedMsg := "testAllergy"
	if resMap["deletedAllergies"] != expectedMsg || resMap["notDeletedAllergies"] != "" {
		t.Errorf("Expected %s; got %s \n", expectedMsg, resMap["deletedAllergies"])
	}
}

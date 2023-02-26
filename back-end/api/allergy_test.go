package main

import (
	"bytes"
	"io/ioutil"

	// "strings"
	"encoding/json"
	// "fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddAllergy(t *testing.T) {
	// Initialize router and connect to DB for this test instance
	InitialUserMigration()
	InitialAllergyMigration()
	InitializeRouter()

	info := Allergy{
		Email:   "js@gmail.com",
		Allergy: "testAllergy",
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

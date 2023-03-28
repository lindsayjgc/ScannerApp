package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateList(t *testing.T) {
	InitialUserMigration()
	InitialListMigration()
	InitializeRouter()

	testTitle := RawTitle{
		Title: "testTitle",
	}

	payload, _ := json.Marshal(testTitle)
	req, _ := http.NewRequest("PUT", "/create-list", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	createTestCookie(req, t)

	rr := httptest.NewRecorder()

	CreateList(rr, req)

	response := rr.Result()
	body, _ := ioutil.ReadAll(response.Body)

	resMap := make(map[string]string)
	json.Unmarshal(body, &resMap)

	if http.StatusCreated != response.StatusCode {
		t.Errorf("Expected %v; got %v\n", http.StatusCreated, response.StatusCode)
	}

	expected := "List successfully created"
	if resMap["message"] != expected {
		t.Errorf("Expected %s; got %s \n", expected, resMap["message"])
	}

	title := GroceryTitle{Email: "unit@test.com", Title: "testTitle"}
	ListDB.Where("email = ? AND title = ?", title.Email, title.Title).Unscoped().Delete(&title)
}

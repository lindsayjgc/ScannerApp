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

func TestAddGroceryItem(t *testing.T) {
	InitialUserMigration()
	InitialListMigration()
	InitializeRouter()

	testItem := RawListItems{
		Title: "testTitle",
		Items: "testitem",
	}

	payload, _ := json.Marshal(testItem)
	req, _ := http.NewRequest("POST", "/add-list-items", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	createTestCookie(req, t)

	rr := httptest.NewRecorder()

	AddGroceryItem(rr, req)

	response := rr.Result()
	body, _ := ioutil.ReadAll(response.Body)

	if http.StatusOK != response.StatusCode {
		t.Errorf("Expected %v; got %v\n", http.StatusOK, response.StatusCode)
	}

	resMap := make(map[string]string)
	json.Unmarshal(body, &resMap)

	expected := "testitem"

	if resMap["addedItems"] != expected || resMap["existingItems"] != "" {
		t.Errorf("Expected %s; got %s \n", expected, resMap["addedItems"])
	}

	item := GroceryItem{Email: "unit@test.com", Title: "testTitle", Item: "testitem"}
	ListDB.Where("email = ? AND title = ? AND item = ?", item.Email, item.Title, item.Item).Unscoped().Delete(&item)
}

func TestDeleteList(t *testing.T) {
	InitialUserMigration()
	InitialListMigration()
	InitializeRouter()

	title := GroceryTitle{Email: "unit@test.com", Title: "testTitle"}
	ListDB.Create(&title)

	testTitle := RawTitles{
		Titles: "testTitle",
	}

	payload, _ := json.Marshal(testTitle)
	req, _ := http.NewRequest("DELETE", "/delete-lists", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	createTestCookie(req, t)

	rr := httptest.NewRecorder()

	DeleteList(rr, req)

	response := rr.Result()
	body, _ := ioutil.ReadAll(response.Body)

	if http.StatusOK != response.StatusCode {
		t.Errorf("Expected %v; got %v\n", http.StatusOK, response.StatusCode)
	}

	resMap := make(map[string]string)
	json.Unmarshal(body, &resMap)

	expectedMsg := "testTitle"

	if resMap["deletedLists"] != expectedMsg || resMap["notDeletedLists"] != "" {
		t.Errorf("Expected %s; got %s \n", expectedMsg, resMap["deletedLists"])
	}
}

func TestDeleteListItem(t *testing.T) {
	InitialUserMigration()
	InitialListMigration()
	InitializeRouter()

	item := GroceryItem{Email: "unit@test.com", Title: "testTitle", Item: "testitem"}
	ListDB.Create(&item)

	testItem := RawListItems{
		Title: "testTitle",
		Items: "testitem",
	}

	payload, _ := json.Marshal(testItem)
	req, _ := http.NewRequest("DELETE", "/delete-list/items", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	createTestCookie(req, t)

	// create ResponseRecorder
	rr := httptest.NewRecorder()

	// invoke target function
	DeleteListItem(rr, req)

	// get the response
	response := rr.Result()
	body, _ := ioutil.ReadAll(response.Body)

	// assertion
	if http.StatusOK != response.StatusCode {
		t.Errorf("Expected %v; got %v\n", http.StatusOK, response.StatusCode)
	}

	resMap := make(map[string]string)
	json.Unmarshal(body, &resMap)

	expectedMsg := "testitem"
	if resMap["deletedItems"] != expectedMsg || resMap["notDeletedItems"] != "" {
		t.Errorf("Expected %s; got %s \n", expectedMsg, resMap["deletedItems"])
	}
}

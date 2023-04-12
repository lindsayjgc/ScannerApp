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

func TestGetLists(t *testing.T) {
	InitialUserMigration()
	InitialListMigration()
	InitializeRouter()

	title1 := GroceryTitle{Email: "unit@test.com", Title: "testTitle1"}
	ListDB.Create(&title1)
	title2 := GroceryTitle{Email: "unit@test.com", Title: "testTitle2"}
	ListDB.Create(&title2)

	req, _ := http.NewRequest("GET", "/get-lists", nil)
	rr := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")

	createTestCookie(req, t)

	GetGroceryTitles(rr, req)

	response := rr.Result()
	body, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		t.Errorf("Error: Expected %v, got %v", http.StatusOK, response.StatusCode)
	}

	resMap := make(map[string]string)
	json.Unmarshal(body, &resMap)

	expectedMsg := "testTitle1,testTitle2"
	if resMap["titles"] != expectedMsg {
		t.Errorf("Expected %s; got %s \n", expectedMsg, resMap["titles"])
	}

	ListDB.Where("email = ? AND title = ?", title1.Email, title1.Title).Unscoped().Delete(&title1)
	ListDB.Where("email = ? AND title = ?", title2.Email, title2.Title).Unscoped().Delete(&title2)
}

func TestGetGroceryList(t *testing.T) {
	InitialUserMigration()
	InitialListMigration()
	InitializeRouter()

	title := GroceryTitle{Email: "unit@test.com", Title: "testTitle"}
	ListDB.Create(&title)
	item1 := GroceryItem{Email: "unit@test.com", Title: "testTitle", Item: "testitem1"}
	ListDB.Create(&item1)
	item2 := GroceryItem{Email: "unit@test.com", Title: "testTitle", Item: "testitem2"}
	ListDB.Create(&item2)

	testTitle := RawTitle{
		Title: "testTitle",
	}

	payload, _ := json.Marshal(testTitle)
	req, _ := http.NewRequest("GET", "/get-list", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	createTestCookie(req, t)

	// create ResponseRecorder
	rr := httptest.NewRecorder()

	// invoke target function
	GetGroceryList(rr, req)

	// get the response
	response := rr.Result()
	body, _ := ioutil.ReadAll(response.Body)

	// assertion
	if http.StatusOK != response.StatusCode {
		t.Errorf("Expected %v; got %v\n", http.StatusOK, response.StatusCode)
	}

	resMap := make(map[string]string)
	json.Unmarshal(body, &resMap)

	expected := `{"title":"testTitle","items":"testitem1,testitem2"}`
	if resMap["title"] != "testTitle" || resMap["items"] != "testitem1,testitem2" {
		t.Errorf("Expected %v; got %v", expected, resMap)
	}

	ListDB.Where("email = ? AND title = ?", title.Email, title.Title).Unscoped().Delete(&title)
	ListDB.Where("email = ? AND title = ? AND item = ?", item1.Email, item1.Title, item1.Item).Unscoped().Delete(&item1)
	ListDB.Where("email = ? AND title = ? AND item = ?", item2.Email, item2.Title, item2.Item).Unscoped().Delete(&item2)

}

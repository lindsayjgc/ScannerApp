package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveQuery(t *testing.T) {
	InitialUserMigration()
	InitialSearchMigration()
	InitializeRouter()

	testQuery := RawSearch{
		Query: "testQuery",
	}

	payload, _ := json.Marshal(testQuery)
	req, _ := http.NewRequest("POST", "/search", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	createTestCookie(req, t)

	rr := httptest.NewRecorder()

	SaveQuery(rr, req)

	response := rr.Result()
	body, _ := ioutil.ReadAll(response.Body)

	resMap := make(map[string]string)
	json.Unmarshal(body, &resMap)

	expected := "Query count updated"
	if resMap["message"] != expected {
		t.Errorf("Expected %s; got %s \n", expected, resMap["message"])
	}

	query := Search{Email: "unit@test.com", Query: "testQuery", Count: 1}
	SearchDB.Where("email = ? AND query = ?", query.Email, query.Query).Unscoped().Delete(&query)
}

func TestGetQueries(t *testing.T) {
	InitialUserMigration()
	InitialSearchMigration()
	InitializeRouter()

	query1 := Search{Email: "unit@test.com", Query: "testQuery1", Count: 1}
	query2 := Search{Email: "unit@test.com", Query: "testQuery2", Count: 2}
	SearchDB.Create(&query1)
	SearchDB.Create(&query2)

	req, _ := http.NewRequest("GET", "search", nil)
	rr := httptest.NewRecorder()
	req.Header.Set("Content-Type", "application/json")

	createTestCookie(req, t)

	GetQueries(rr, req)

	response := rr.Result()
	body, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		t.Errorf("Error: Expected %v, got %v", http.StatusOK, response.StatusCode)
	}

	var resList []Query
	json.Unmarshal(body, &resList)

	expected := []string{"testQuery1", "testQuery2"}
	if resList[0].Query != "testQuery1" {
		t.Errorf("Expected %s; got %s \n", expected[0], resList[0].Query)
	}
	if resList[1].Query != "testQuery2" {
		t.Errorf("Expected %s; got %s \n", expected[1], resList[1].Query)
	}

	SearchDB.Where("email = ? AND query = ?", query1.Email, query1.Query).Delete(&query1)
	SearchDB.Where("email = ? AND query = ?", query2.Email, query2.Query).Delete(&query2)

}

func TestDeleteQuery(t *testing.T) {
	InitialUserMigration()
	InitialSearchMigration()
	InitializeRouter()

	query := Search{Email: "unit@test.com", Query: "testQuery", Count: 1}
	SearchDB.Create(&query)

	testQuery := RawSearch{
		Query: "testQuery",
	}

	payload, _ := json.Marshal(testQuery)
	req, _ := http.NewRequest("DELETE", "/search", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	createTestCookie(req, t)

	rr := httptest.NewRecorder()

	RemoveQuery(rr, req)

	response := rr.Result()
	body, _ := ioutil.ReadAll(response.Body)

	if http.StatusOK != response.StatusCode {
		t.Errorf("Expected %v; got %v\n", http.StatusOK, response.StatusCode)
	}

	resMap := make(map[string]string)
	json.Unmarshal(body, &resMap)

	expected := "Query successfully deleted"
	if resMap["message"] != expected {
		t.Errorf("Expected %s; got %s \n", expected, resMap["message"])
	}
}

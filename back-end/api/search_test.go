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

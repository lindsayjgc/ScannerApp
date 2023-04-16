package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetLabels(t *testing.T) {
	InitialLabelMigration()
	InitializeRouter()

	req, _ := http.NewRequest("GET", "/api/label?type=diet", nil)
	req.Header.Set("Content-Type", "application/json")

	// Create a mock request recorder
	rr := httptest.NewRecorder()

	// Send the request and recorder
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned the wrong status: got %v, expected %v", status, http.StatusCreated)
	}

	expected := `{"labels":"balanced,high-fiber,high-protein,low-carb,low-fat,low-sodium"}`

	// Remove any linebreaks from the response writer body
	body := strings.Replace(rr.Body.String(), "\n", "", -1)
	body = strings.Replace(body, "\r", "", -1)

	if body != expected {
		t.Errorf("Handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected)
	}
}
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAddFavorite(t *testing.T) {
	InitialUserMigration()
	InitialFavoriteMigration()
	InitializeRouter()

	favorite := RawFavorite{
		Favorite: "Test",
		Code: "testcode",
		Image: "testimagelink.com",
	}
	
	// Create a mock request
	payload, _ := json.Marshal(favorite)
	req, _ := http.NewRequest("POST", "/api/favorite", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	
	// Create test cookie to simulate user login
	createTestCookie(req, t)

	// Create a mock request recorder
	rr := httptest.NewRecorder()

	// Send the request and recorder
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned the wrong status: got %v, expected %v", status, http.StatusCreated)
	}

	expected := `{"message":"Product successfully favorited"}`

	// Remove any linebreaks from the response writer body
	body := strings.Replace(rr.Body.String(), "\n", "", -1)
	body = strings.Replace(body, "\r", "", -1)

	if body != expected {
		t.Errorf("Handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected)
	}

	// Delete the favorite that was created for the test
	var deleteFavorite Favorite
	FavoriteDB.Where("email = ?", "unit@test.com").Where("code = ?", "testcode").Unscoped().Delete(&deleteFavorite)
}
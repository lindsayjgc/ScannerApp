package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetFavorites(t *testing.T) {
	InitialUserMigration()
	InitialFavoriteMigration()
	InitializeRouter()

	// Create two test favorites to be fetched
	testFavorite1 := Favorite{
		Email: "unit@test.com",
		Favorite: "Test1",
		Code: "testcode1",
		Image: "testimagelink.com",
	}

	testFavorite2 := Favorite{
		Email: "unit@test.com",
		Favorite: "Test2",
		Code: "testcode2",
		Image: "testimagelink.com",
	}

	// Add test favorites to the DB
	FavoriteDB.Create(&testFavorite1)
	FavoriteDB.Create(&testFavorite2)

	// Create a mock request
	req, _ := http.NewRequest("GET", "/api/favorite", nil)
	req.Header.Set("Content-Type", "application/json")
	
	// Create test cookie to simulate user login
	createTestCookie(req, t)

	// Create a mock request recorder
	rr := httptest.NewRecorder()

	// Send the request and recorder
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned the wrong status: got %v, expected %v", status, http.StatusCreated)
	}

	expected := `[{"favorite":"Test1","code":"testcode1","image":"testimagelink.com"},{"favorite":"Test2","code":"testcode2","image":"testimagelink.com"}]`
	body := strings.Replace(rr.Body.String(), "\n", "", -1);
	body = strings.Replace(body, "\r", "", -1);

	if body != expected {
		t.Errorf("Handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected);
	}

	// Clean up test data from DB
	var deletedFavorites []RawFavorite
	FavoriteDB.Table("favorites").Where("email = ?", "unit@test.com").Unscoped().Delete(&deletedFavorites)
}

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

func TestDeleteFavorite(t *testing.T) {
	InitialUserMigration()
	InitialFavoriteMigration()
	InitializeRouter()

	// Create mock favorite to be delete
	favorite := Favorite{
		Email: "unit@test.com",
		Favorite: "Test",
		Code: "testcode",
		Image: "testimagelink.com",
	}
	FavoriteDB.Create(&favorite)
	
	// Create a mock request
	payload, _ := json.Marshal(favorite)
	req, _ := http.NewRequest("DELETE", "/api/favorite", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	
	// Create test cookie to simulate user login
	createTestCookie(req, t)

	// Create a mock request recorder
	rr := httptest.NewRecorder()

	// Send the request and recorder
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned the wrong status: got %v, expected %v", status, http.StatusCreated)
	}

	expected := `{"message":"Favorite successfully deleted"}`

	// Remove any linebreaks from the response writer body
	body := strings.Replace(rr.Body.String(), "\n", "", -1)
	body = strings.Replace(body, "\r", "", -1)

	if body != expected {
		t.Errorf("Handler returned unexpected body: got %v, expected %v", rr.Body.String(), expected)
	}

	// Delete the favorite with GORM to remove the soft-deleted
	var deleteFavorite Favorite
	FavoriteDB.Where("email = ?", "unit@test.com").Where("code = ?", "testcode").Unscoped().Delete(&deleteFavorite)
}
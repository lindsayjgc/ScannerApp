package main

import (
	"testing"
)

func TestFetchFoodListPage(t *testing.T) {
	// Ensure that GetFoodList returns a non-empty list
	foodList, err := fetchFoodListPage(1)
	if err != nil {
		t.Errorf("Error fetching food list: %v", err)
	}
	if len(foodList) == 0 {
		t.Errorf("Food list is empty")
	}
}

func TestGetAllNutrients(t *testing.T) {
	// Call the function being tested
	err := GetAllNutrients()

	// Check that there were no errors
	if err != nil {
		t.Errorf("GetAllNutrients() returned an error: %v", err)
	}

	// Check that nutrient data was populated for at least one food
	found := false
	for _, food := range foodList {
		if food.Nutrients != nil {
			found = true
			break
		}
	}
	if !found {
		t.Error("GetAllNutrients() did not populate nutrient data for any food")
	}
}

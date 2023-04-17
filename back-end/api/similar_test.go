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

func TestGetFoodList(t *testing.T) {
	// Ensure that GetFoodList returns a non-empty list
	foodList, err := GetFoodList()
	if err != nil {
		t.Errorf("Error fetching food list: %v", err)
	}
	if len(foodList) == 0 {
		t.Errorf("Food list is empty")
	}
}

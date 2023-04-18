package main

import (
	"math"
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

func TestCosineSimilarity(t *testing.T) {
	// Test the CosineSimilarity function with two identical nutrient maps
	x := map[string]float64{
		"protein": 10,
		"fat":     20,
		"carb":    30,
	}
	y := map[string]float64{
		"protein": 10,
		"fat":     20,
		"carb":    30,
	}
	similarity := CosineSimilarity(x, y)
	similarity = math.Round(similarity)
	if similarity != 1 {
		t.Errorf("Cosine similarity of identical maps should be 1, got %f", similarity)
	}

	// Test the CosineSimilarity function with two completely different nutrient maps
	x = map[string]float64{
		"protein": 10,
		"fat":     20,
		"carb":    30,
	}
	y = map[string]float64{
		"calories": 100,
		"sodium":   200,
	}
	similarity = CosineSimilarity(x, y)
	if similarity != 0 {
		t.Errorf("Cosine similarity of completely different maps should be 0, got %f", similarity)
	}

	// Test the CosineSimilarity function with two partially overlapping nutrient maps
	x = map[string]float64{
		"protein": 10,
		"fat":     20,
		"carb":    30,
	}
	y = map[string]float64{
		"protein": 10,
		"fat":     20,
		"carb":    40,
	}
	similarity = CosineSimilarity(x, y)
	similarity = math.Trunc(similarity*100) / 100
	if similarity != 0.99 {
		t.Errorf("Cosine similarity of partially overlapping maps should be 0.99, got %f", similarity)
	}
}

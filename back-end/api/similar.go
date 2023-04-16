package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Food struct {
	FdcID        int64
	Description  string
	FoodCategory string
	DataPoints   int64
	Nutrients    map[string]float64
}

type FoodList struct {
	Foods []Food
}

func fetchFoodListPage(pageNumber int) ([]Food, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	response, err := http.Get(fmt.Sprintf("https://api.nal.usda.gov/fdc/v1/foods/list?dataType=Foundation&pageSize=200&pageNumber=%d&api_key=%s", pageNumber, os.Getenv("API_KEY")))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var foodList []Food
	err = json.Unmarshal(body, &foodList)
	if err != nil {
		return nil, err
	}

	return foodList, nil
}

func GetFoodList() ([]Food, error) {
	foodList, err := fetchFoodListPage(1)
	if err != nil {
		return nil, err
	}

	foodListEnd, err := fetchFoodListPage(2)
	if err != nil {
		return nil, err
	}

	// Append foodListEnd to foodList
	foodList = append(foodList, foodListEnd...)
	fmt.Println(foodList)

	return foodList, nil
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"

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

type FoodNutrient struct {
	Number                string  `json:"number"`
	Name                  string  `json:"name"`
	Amount                float64 `json:"amount"`
	UnitName              string  `json:"unitName"`
	DerivationCode        string  `json:"derivationCode"`
	DerivationDescription string  `json:"derivationDescription"`
}

type NewFood struct {
	FdcID           int64          `json:"fdcId"`
	Description     string         `json:"description"`
	DataType        string         `json:"dataType"`
	PublicationDate string         `json:"publicationDate"`
	NdbNumber       string         `json:"ndbNumber"`
	FoodNutrients   []FoodNutrient `json:"foodNutrients"`
}

type SimilarityResult struct {
	FdcID      int64
	Similarity float64
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

	return foodList, nil
}

func GetAllNutrients() error {
	foodList, err = GetFoodList()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	var wg sync.WaitGroup
	nutrientErrors := make(chan error, len(foodList))

	for i, food := range foodList {
		wg.Add(1)
		go func(i int, food Food) {
			defer wg.Done()
			nutrients, err := GetFoodNutrients(food.FdcID)
			if err != nil {
				nutrientErrors <- fmt.Errorf("Error fetching nutrients for food with FdcID %d: %v", food.FdcID, err)
				return
			}
			foodList[i].Nutrients = nutrients
		}(i, food)
	}

	wg.Wait()
	close(nutrientErrors)

	for err := range nutrientErrors {
		fmt.Println(err)
	}

	return nil
}

func GetFoodNutrients(fdcID int64) (map[string]float64, error) {
	fmt.Printf("Fetching nutrients for FdcID %d...\n", fdcID)
	nutrientIds := []int{203, 204, 205}
	nutrientChunks := chunkNutrientIds(nutrientIds, 3)

	nutrients := make(map[string]float64)
	var mutex sync.Mutex
	var wg sync.WaitGroup

	errChan := make(chan error, len(nutrientChunks))

	for _, chunk := range nutrientChunks {
		wg.Add(1)
		go func(chunk []int) {
			defer wg.Done()

			url := fmt.Sprintf("https://api.nal.usda.gov/fdc/v1/food/%d?format=abridged&api_key=%s", fdcID, os.Getenv("API_KEY"))
			for _, nutrient := range chunk {
				url += fmt.Sprintf("&nutrients=%d", nutrient)
			}

			response, err := http.Get(url)
			if err != nil {
				errChan <- err
				return
			}
			defer response.Body.Close()

			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				errChan <- err
				return
			}

			var food NewFood
			err = json.Unmarshal(body, &food)
			if err != nil {
				errChan <- err
				return
			}

			mutex.Lock()
			for _, nutrient := range food.FoodNutrients {
				nutrients[nutrient.Name] = nutrient.Amount
			}
			mutex.Unlock()

		}(chunk)
	}

	wg.Wait()
	close(errChan)

	for err := range errChan {
		return nil, err
	}

	return nutrients, nil
}

func chunkNutrientIds(nutrientIds []int, chunkSize int) [][]int {
	var chunks [][]int
	for i := 0; i < len(nutrientIds); i += chunkSize {
		end := i + chunkSize

		if end > len(nutrientIds) {
			end = len(nutrientIds)
		}

		chunks = append(chunks, nutrientIds[i:end])
	}
	return chunks
}

func CosineSimilarity(x, y map[string]float64) float64 {
	// Get the intersection of keys between the two maps
	var keys []string
	for key := range x {
		if _, ok := y[key]; ok {
			keys = append(keys, key)
		}
	}

	// If there are no common keys, return 0
	if len(keys) == 0 {
		return 0
	}

	// Compute the dot product and the magnitudes
	var dotProduct, xMagnitude, yMagnitude float64
	for _, key := range keys {
		dotProduct += x[key] * y[key]
		xMagnitude += x[key] * x[key]
		yMagnitude += y[key] * y[key]
	}

	// Compute the cosine similarity
	return dotProduct / (math.Sqrt(xMagnitude) * math.Sqrt(yMagnitude))
}

func GetSimilarFoods(search string, foodList []Food) []Food {

	// Get the search food
	var searchFood Food
	for _, food := range foodList {
		if strings.Contains(strings.ToLower(food.Description), strings.ToLower(search)) {
			searchFood = food
			break
		}
	}
	if searchFood.FdcID == 0 {
		fmt.Println("Search food not found")
		return nil
	}

	// return similarFoods
	// Compute the cosine similarity for each food in the list
	var results []SimilarityResult
	for _, food := range foodList {
		if food.FdcID != searchFood.FdcID {
			similarity := CosineSimilarity(searchFood.Nutrients, food.Nutrients)
			result := SimilarityResult{
				FdcID:      food.FdcID,
				Similarity: similarity,
			}
			results = append(results, result)
		}
	}

	// Sort the results by similarity in descending order
	sort.Slice(results, func(i, j int) bool {
		return results[i].Similarity > results[j].Similarity
	})

	// Get the most similar foods based on the coefficient
	var similarFoods []Food
	for _, result := range results {
		for _, food := range foodList {
			if food.FdcID == result.FdcID {
				similarFoods = append(similarFoods, food)
				break
			}
		}
		if len(similarFoods) >= 5 {
			break
		}
	}

	return similarFoods
}

var foodList []Food

func SimilarFoods(w http.ResponseWriter, r *http.Request) {
	if len(foodList) == 0 {
		GetAllNutrients()
	}

	var newSearch RawSearch
	err = json.NewDecoder(r.Body).Decode(&newSearch)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON"))
		return
	}

	search := newSearch.Query
	fmt.Print(search)

	similarFoods := GetSimilarFoods(search, foodList)

	res := make(map[int64]string)
	for _, food := range similarFoods {
		res[food.FdcID] = food.Description
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

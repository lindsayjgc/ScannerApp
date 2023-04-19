package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var ListDB *gorm.DB

type GroceryItem struct {
	gorm.Model
	Email string `json:"email"`
	Title string `json:"title"`
	Item  string `json:"item"`
}

type GroceryTitle struct {
	gorm.Model
	Email string `json:"email"`
	Title string `json:"title"`
}

type RawListItems struct {
	Title string `json:"title"`
	Items string `json:"items"`
}

type RawTitles struct {
	Titles string `json:"titles"`
}

type RawTitle struct {
	Title string `json:"title"`
}

func InitialListMigration() {
	ListDB, err = gorm.Open(sqlite.Open(DB_PATH), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("Error connecting to DB.")
	}

	err := godotenv.Load()

	if err != nil {
		fmt.Println(err)
		panic("Error loading .env file.")
	}

	jwtKey = []byte(os.Getenv("SECRET_KEY"))

	ListDB.AutoMigrate(&GroceryItem{})
	ListDB.AutoMigrate(&GroceryTitle{})
}

func CreateList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var listTitle RawTitle
	err = json.NewDecoder(r.Body).Decode(&listTitle)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON"))
		return
	}

	// Check if title already exists in database
	var existingTitle GroceryTitle
	res := ListDB.First(&existingTitle, "email = ? AND title = ?", claims.Email, listTitle.Title)

	if res.Error == nil {
		// title already exists in db
		w.WriteHeader(http.StatusAlreadyReported)
		json.NewEncoder(w).Encode(GenerateResponse("List already exists"))
		return
	}

	title := GroceryTitle{Title: listTitle.Title, Email: claims.Email}
	ListDB.Create(&title)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(GenerateResponse("List successfully created"))

}

func AddGroceryItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Check for logged in user and get their email
	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var rawlistItems RawListItems
	err = json.NewDecoder(r.Body).Decode(&rawlistItems)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON"))
		return
	}

	newItems := strings.Split(strings.ToLower(string(rawlistItems.Items)), ",")

	// Check if item is already on list
	var existingItemsSlice []string
	ListDB.Model(GroceryItem{}).Where("email = ? AND title = ?", claims.Email, rawlistItems.Title).Select("item").Find(&existingItemsSlice)

	// Convert the slice of items into a map for efficiency later
	existingItems := make(map[string]bool)
	for _, v := range existingItemsSlice {
		existingItems[string(v)] = true
	}

	// Check for existing items and add them to DB if not
	var addedItems []string
	var notAddedItems []string
	for _, v := range newItems {
		if !existingItems[v] { // If this new item does not already exist
			item := GroceryItem{Email: claims.Email, Title: rawlistItems.Title, Item: v}
			addedItems = append(addedItems, v)
			ListDB.Create(&item)
		} else { // Otherwise, add to a list of already existing items
			notAddedItems = append(notAddedItems, v)
		}
	}

	res := make(map[string]string)
	res["addedItems"] = strings.Join(addedItems, ",")
	res["existingItems"] = strings.Join(notAddedItems, ",")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// function for list of titles
func GetGroceryTitles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var groceryTitlesSlice []string
	ListDB.Model(GroceryTitle{}).Where("email = ?", claims.Email).Select("title").Find(&groceryTitlesSlice)

	var groceryTitles RawTitles
	if len(groceryTitlesSlice) == 0 {
		groceryTitles.Titles = "NONE"
	} else {
		groceryTitles.Titles = strings.Join(groceryTitlesSlice, ",")
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(groceryTitles)
}

// function for returning list and all items on it
func GetGroceryList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var listTitle RawTitle
	err = json.NewDecoder(r.Body).Decode(&listTitle)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON"))
		return
	}

	var itemsSlice []string
	ListDB.Model(GroceryItem{}).Where("email = ? AND title = ?", claims.Email, listTitle.Title).Select("item").Find(&itemsSlice)

	var list RawListItems
	list.Title = listTitle.Title
	if len(itemsSlice) == 0 {
		list.Items = "NONE"
	} else {
		list.Items = strings.Join(itemsSlice, ",")
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(list)

}

func DeleteList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var rawListsToDelete RawTitles
	err = json.NewDecoder(r.Body).Decode(&rawListsToDelete)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON"))
		return
	}

	listsToDelete := strings.Split(string(rawListsToDelete.Titles), ",")

	var existingListsSlice []string
	ListDB.Model(GroceryTitle{}).Where("email = ?", claims.Email).Select("title").Find(&existingListsSlice)

	// Convert the slice of lists into a map for efficiency later
	existingLists := make(map[string]bool)
	for _, v := range existingListsSlice {
		existingLists[string(v)] = true
	}

	var deletedLists []string
	var notDeletedLists []string
	for _, v := range listsToDelete {
		if existingLists[v] {
			deletedLists = append(deletedLists, v)
			title := GroceryTitle{Email: claims.Email, Title: v}
			ListDB.Where("email = ? AND title = ?", title.Email, title.Title).Delete(&title)
			ListDB.Where("email = ? AND title = ?", title.Email, title.Title).Delete(&GroceryItem{})
		} else {
			notDeletedLists = append(notDeletedLists, v)
		}
	}

	res := make(map[string]string)
	res["deletedLists"] = strings.Join(deletedLists, ",")
	res["notDeletedLists"] = strings.Join(notDeletedLists, ",")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func DeleteListItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, err, resStatus := CheckCookie(w, r)

	if err != nil {
		w.WriteHeader(resStatus)
		json.NewEncoder(w).Encode(GenerateResponse(err.Error()))
		return
	}

	var rawItemsToDelete RawListItems
	err = json.NewDecoder(r.Body).Decode(&rawItemsToDelete)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(GenerateResponse("Error decoding JSON"))
		return
	}

	itemsToDelete := strings.Split(string(rawItemsToDelete.Items), ",")

	var existingItemsSlice []string
	ListDB.Model(GroceryItem{}).Where("email = ? AND title = ?", claims.Email, rawItemsToDelete.Title).Select("item").Find(&existingItemsSlice)

	// Convert the slice of items into a map for efficiency later
	existingItems := make(map[string]bool)
	for _, v := range existingItemsSlice {
		existingItems[string(v)] = true
	}

	var deletedItems []string
	var notDeletedItems []string
	for _, v := range itemsToDelete {
		if existingItems[v] {
			deletedItems = append(deletedItems, v)
			item := GroceryItem{Email: claims.Email, Title: rawItemsToDelete.Title, Item: v}
			ListDB.Where("email = ? AND title = ? AND item = ?", item.Email, item.Title, item.Item).Delete(&item)
		} else {
			notDeletedItems = append(notDeletedItems, v)
		}
	}

	res := make(map[string]string)
	res["deletedItems"] = strings.Join(deletedItems, ",")
	res["notDeletedItems"] = strings.Join(notDeletedItems, ",")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

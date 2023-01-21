package openmensa

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var (
	Base_URL = "https://openmensa.org/api/v2/"
)

type Canteen struct {
	ID          int           `json:"id"`
	Name        string        `json:"name"`
	City        string        `json:"city"`
	Address     string        `json:"address"`
	Coordinates []interface{} `json:"coordinates"`
	Notes       []string      `json:"notes"`
}

type Meal struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Category string    `json:"category"`
	Prices   MealPrice `json:"prices"`
}

type MealPrice struct {
	Students  float64 `json:"students"`
	Employees float64 `json:"employees"`
	Pupils    float64 `json:"pupils"`
	Others    float64 `json:"others"`
}

// json to canteen

// get all canteens from openmensa API
func GetCanteens() []Canteen {
	req, err := http.NewRequest("GET", Base_URL+"canteens", nil)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer res.Body.Close()

	print(res.StatusCode)

	// parse json
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	var canteens []Canteen

	json.Unmarshal(body, &canteens)

	return canteens
}

// get meals for canteen on specific date
func GetMeals(canteenID int, date string) []Meal {
	id := strconv.Itoa(canteenID)

	req, err := http.NewRequest("GET", Base_URL+"canteens/"+id+"/days/"+date+"/meals", nil)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer res.Body.Close()

	print(res.StatusCode)

	// parse json
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	var meals []Meal

	json.Unmarshal(body, &meals)

	return meals
}

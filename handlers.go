package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Index handle requests to "/"
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to carboservice: GET /carbs/01009 for nutrition info about cheddar cheese")
}

// CarbsIndex handle requests to "/carbs"
func CarbsIndex(w http.ResponseWriter, r *http.Request) {
	carbs := Carbs{
		Carb{Name: "Write carboservice"},
		Carb{Name: "Read about CodePipeline"},
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(carbs); err != nil {
		panic(err)
	}
}

// CarbsOne handle requests to "/carbs/:id"
func CarbsOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ndbID := vars["ndbID"]
	url := nutrientURL + "&ndbno=" + ndbID
	log.Println(url)
	resp, err := http.Get(url)

	if err != nil {
		log.Println(float64(resp.StatusCode) / 200)
		panic(err)
	}

	if float64(resp.StatusCode)/200 > 1.495 {
		http.Error(w, "Could not connect to Nutrient DB", http.StatusInternalServerError)
		return
	}

	var n NutrientAPIResponse
	json.NewDecoder(resp.Body).Decode(&n)
	if err != nil {
		panic(err)
	}

	var response = n.transform()
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}

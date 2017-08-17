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
	fmt.Fprintf(w, "Hello Friends!")
}

// CarbsIndex handle requests to "/carbs"
func CarbsIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Usage: GET /carbs/:id for nutrition info")
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

package main

import (
	"log"
	"net/http"
	"os"
)

var nutrientURL = "http://api.nal.usda.gov/ndb/reports/?type=b&format=json&api_key=" + os.Getenv("APIKEY")

func main() {
	router := NewRouter()

	log.Println("starting carboservice on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

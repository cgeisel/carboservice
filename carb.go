package main

import (
	"fmt"
	"strconv"
)

// Carb contains food Name, CarbFactor (multiply per 1g for total carbs), PerChoice (amount of food with 15g carbs)
type Carb struct {
	Name       string  `json:"name"`
	CarbFactor float64 `json:"carb-factor"`
	PerChoice  float64 `json:"per-choice"`
}

// Carbs slice of Carb structs
type Carbs []Carb

type Nutrient struct {
	NutrientID int     `json:"nutrient_id,string"`
	Value      float64 `json:"value,string"`
}

type NutrientAPIResponse struct {
	Report struct {
		Food struct {
			Ndbno     int    `json:"ndbno,string"`
			Name      string `json:"name"`
			Nutrients []Nutrient
		}
	}
}

type CarbAPIResponse struct {
	ID         int
	Name       string
	CarbFactor float64
	PerChoice  int
}

func (n NutrientAPIResponse) transform() CarbAPIResponse {
	var response CarbAPIResponse
	response.Name = n.Report.Food.Name
	response.ID = n.Report.Food.Ndbno
	for _, v := range n.Report.Food.Nutrients {
		if v.NutrientID == 205 {
			var carbfactor, err = strconv.ParseFloat(fmt.Sprintf("%.2f", v.Value/100), 64)
			if err != nil {
				panic(err)
			}
			response.CarbFactor = carbfactor
			response.PerChoice = int(v.Value * 15)
		}
	}
	return response
}

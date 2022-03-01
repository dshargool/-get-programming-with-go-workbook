package main

import (
	"encoding/json"
	"fmt"
)

type location struct {
	lat  float64
	long float64
}

func main() {

	// Define a struct as a var
	var opportunity struct {
		lat  float64
		long float64
	}

	opportunity.lat = 233.3
	opportunity.long = -2222.2

	// Use the location struct type to define the variable
	var curiousity location
	curiousity.lat = 10.2
	curiousity.long = 543.2
	var spirit location // Re-use struct as a type
	spirit.lat = 12.2
	spirit.long = -322.2

	fmt.Println(opportunity)
	fmt.Println(curiousity)
	fmt.Println(spirit)

	// Initialize with composite literals
	opportunity = location{lat: 33.3, long: -99.9}
	fmt.Println(opportunity)
	// Initialize with composite literals without name.  Order matters here!
	opportunity = location{1.1, 2.2}
	fmt.Println(opportunity)

	// Copying structures
	obj_one := location{-4.5, 137.4}
	obj_two := obj_one
	obj_two.lat += 0.111
	fmt.Println(obj_one, obj_two)

	// This is bad!
	lats := []float64{-4.5, -14.56, -1.9}
	longs := []float64{137.4, 175.47, 354.47}
	fmt.Println(lats, longs)

	type named_location struct {
		Name string
		Lat  float64
		Long float64
	}

	locations := []named_location{
		{Name: "Brad", Lat: 4.44, Long: 6.66},
		{Name: "Colum", Lat: 1.22, Long: 99.99},
		{Name: "Chall", Lat: -1.11, Long: -0.22},
	}

	fmt.Println(locations)

	// Encoding to JSON
	type json_location struct {
		Lat, Long float64
	}

	curiousity_json := json_location{-4.58, 137.44}
	bytes, _ := json.Marshal(curiousity_json)
	fmt.Println(string(bytes))

	type json_tags_location struct {
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	curiousity_tags_json := json_tags_location{-4.58, 137.44}
	bytes, _ = json.Marshal(curiousity_tags_json)
	fmt.Println(string(bytes))

	// Landing.go
	bytes, _ = json.MarshalIndent(locations, "", "    ")
	fmt.Println(string(bytes))
}

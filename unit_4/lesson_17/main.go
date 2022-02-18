package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	terrestrial := planets[0:4] // Indexes 0 to 3
	gasGiants := planets[4:6]   // Indexes 4 to 5
	iceGiants := planets[6:8]   // Indexes 6 to 7
	fmt.Println(terrestrial)
	fmt.Println(gasGiants)

	iceGiants[0] = "Tootoo"
	fmt.Println(iceGiants)
	fmt.Println(planets)
	planet_slice := planets[:]

	// Convert the planet_splice to the type StringSlice which is part of the sort package and then sort it
	sort.StringSlice(planet_slice).Sort()
	fmt.Println(planet_slice)
}

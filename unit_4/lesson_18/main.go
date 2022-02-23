package main

import "fmt"

func terraform(prefix string, worlds ...string) []string {
	newWorlds := make([]string, len(worlds)) // Makes a new slice instead of modifying directly
	for i := range worlds {
		newWorlds[i] = prefix + " " + worlds[i]
	}
	return newWorlds
}

func main() {

	dwarfs := []string{"One", "Two", "Three", "Four", "Five"}
	fmt.Println(len(dwarfs))      // Outputs 5
	fmt.Println(cap(dwarfs))      // Outputs 5
	fmt.Println(len(dwarfs[1:2])) // Outputs 1
	fmt.Println(cap(dwarfs[1:2])) // Outputs 4

	dwarfs = []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Println(dwarfs)
	dwarfs = append(dwarfs, "Orcus")
	fmt.Println(dwarfs)
	// See that the append function is variadic
	dwarfs = append(dwarfs, "Salacia", "Quaoar", "Sedna")
	fmt.Println(dwarfs)

	planets := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}
	terrestrial := planets[0:4:4] // Length 4, Capacity 4
	worlds := append(terrestrial, "Ceres")
	fmt.Println(planets)

	terrestrial = planets[0:4]
	worlds = append(terrestrial, "Ceres")
	fmt.Println(worlds)
	fmt.Println(planets) // Overwrites the 5th element in the underlying array

	twoPlanets := terraform("New", "Venus", "Mars") // Notice how to pass the slice as multiple arguments
	fmt.Println(twoPlanets)

	my_planets := []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("New", my_planets...) // Notice how to pass the slice as multiple arguments
	fmt.Println(newPlanets)

	// capacity.go
	myBigArray := []int{1, 2, 3, 4}
	for i := 0; i < 1000; i++ {
		myBigArray = append(myBigArray, i)
		//fmt.Println(cap(myBigArray))
	}

}

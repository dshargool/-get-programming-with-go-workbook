package main

import (
	"fmt"
	"math/rand"
)

// Scope inside the whole package
// Short declaration is not available here
var era = "AD"

func scope() {
	// Function scope is here
	var count = 0
	for count < 10 {
		// New scope starts here with the curly braces
		//		var num = rand.Intn(10) + 1
		num := count + 1
		fmt.Println(num)
		count++
		// New scope ends here
	}

	// Short declaration uses :=
	shortcount := 3
	fmt.Println(shortcount)

	// Declaration in for loop
	for count = 10; count > 0; count-- {
		fmt.Println(count)
	}

	// Declaration in an if statement
	if num_space := rand.Intn(3); num_space == 0 {
		fmt.Println("Space Adventures")
	} else if num_space == 1 {
		fmt.Println("SpaceX")
	} else {
		fmt.Println("Virgin Galactic")
		fmt.Println(num_space)
	}

	year := 2022

	// Month is scoped to the switch statement
	switch month := rand.Intn(12) + 1; month {
	case 2:
		day := rand.Intn(28) + 1
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1
		fmt.Println(era, year, month, day)
	}
}

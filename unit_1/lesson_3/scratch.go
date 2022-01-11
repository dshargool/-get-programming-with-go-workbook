package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	// True or false
	// In Go only 'true' and 'false' are valid booleans (no 0 or "")
	fmt.Println("You find yourself in a dimly lit cavern.")
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	//var wearShades = true
	//var read = strings.Contains(command, "read")
	fmt.Println("You leave the cave:", exit)

	// Comparisons
	fmt.Println("There is a sign at the entrance that reads 'No Minors'.")
	var age = 41
	var minor = age < 18
	fmt.Printf("At age %v, am I a minor? %v\n", age, minor)
	fmt.Println("apple" > "banana")

	// If-statements
	command = "go east"
	if command == "go east" {
		fmt.Println("You head further up the mountain")
	} else if command == "go inside" {
		fmt.Println("You enter the cave where you live out the rest of your life.")
	} else {
		fmt.Println("Didn't quite get that.")
	}

	var room = "mountain"
	if room == "cave" {
		fmt.Println("You are in a deep, moist, cavern")
	} else if room == "entrance" {
		fmt.Println("You are on the precipice of entry")
	} else if room == "mountain" {
		fmt.Println("You are on the mountain.  You can move it brick by brick.")
	} else {
		fmt.Println("Where are you?")
	}

	// Use and/or with modulus
	// Calculate leap years
	var year = 2000
	fmt.Printf("Is the year %v a leap year?\n", year)
	var leap = year%400 == 0 || year%4 == 0 && year%100 != 0
	if leap {
		fmt.Println("It is a leap year!\n")
	} else {
		fmt.Println("It isn't leap year!\n")

	}

	// Switch statements
	room = "lake"

	switch {
	case room == "cave":
		fmt.Println("You are in the cave of caves")
	case room == "lake":
		fmt.Println("You are in the lake")
		fallthrough // This falls through to the next case (same as in C)
	case room == "underwater":
		fmt.Println("You are under the sea")
	}
	// Quick Check 3.5
	switch room {
	case "cave":
		fmt.Println("You are in the cave of caves")
	case "lake":
		fmt.Println("You are in the lake")
	case "underwater":
		fmt.Println("You are under the sea")
	}

	// Looping
	var count = 100
	for count > 0 {
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
		if rand.Int31n(99) == 0 {
			break
		}
	}
	if count == 0 {
		fmt.Println("Liftoff!")
	} else {
		fmt.Println("Failure to launch")

	}

}

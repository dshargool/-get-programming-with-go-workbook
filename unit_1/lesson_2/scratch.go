package main

import (
	"fmt"
	"math/rand"
)

func scratch() {
	fmt.Printf("My weight on the surface of Mars is %v lbs and I would be %v years old.\n", 149*0.3783, 41*365/687)
	// Negative number pads to the right.  Positive pads to the left
	fmt.Printf("%19v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-19v $%4v\n", "Virgin Galactic", 100)

	// 2.4.2-2.4.3 Formatted Print and Constants and Variables

	const lightSpeed = 299792                  // km/s
	const hoursPerDay, minutesPerHour = 24, 60 // set up time constants
	const spaceXspeed = 100800                 // km/h
	var distance = 56000000                    // km

	fmt.Println(distance/lightSpeed, "seconds")
	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")
	distance = 96300000
	fmt.Println(distance/spaceXspeed*minutesPerHour, "minutes")

	// 2.4.2 Increment and Assignment Operators
	var weight = 149.0
	// weight = weight * 0.3783 is the same as
	weight *= 0.3783

	// Increment/Decrement has a special ++ or -- operator
	var age = 41
	age = age + 1
	fmt.Println(age)
	age += 1
	fmt.Println(age)
	age++
	fmt.Println(age)

	// Quickcheck 2.5
	fmt.Println(weight)
	weight -= 2
	fmt.Println(weight)

	// Random number generation
	var num = rand.Intn(10) + 1
	fmt.Println(num)
	num = rand.Intn(10) + 1
	fmt.Println(num)

	// Quick Check 2.6
	distance = rand.Intn(401000000-56000000) + 56000000
	fmt.Println(distance)

}

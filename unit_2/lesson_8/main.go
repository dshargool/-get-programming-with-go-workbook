package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Distance to alpha centauri
	var distance int64 = 41.3e12
	fmt.Println(distance)

	// Distance to andromeda which overflows uint64
	// distance = 24e18

	// Can use the big package with 3 types:
	// big.Int
	// big.Float
	// big.Rat (for fractions like 1/3)

	// To use big.Int for a large number
	big_distance := new(big.Int)
	// Set to 24 quintillion with base 10
	big_distance.SetString("24000000000000000000", 10)
	fmt.Println(big_distance)

	// Two ways
	big_dist := big.NewInt(3000)
	fmt.Println(big_dist)
	// or
	super_big_dist := new(big.Int)
	super_big_dist.SetString("9001", 10)
	fmt.Println(super_big_dist)

	fmt.Println("Start on Lesson 6")

	// Constants
	// Untyped and are dealth with using the `big` package at compile time.
	// This does not mean that we can assign them to a regular variable and use them.

	const major_dwarf_dist_km = 236000000000000000
	const seconds_per_year = 60 * 60 * 24 * 365
	const lightspeed_ms = 299792
	lightyears := major_dwarf_dist_km / lightspeed_ms / seconds_per_year
	fmt.Println(lightyears)
}

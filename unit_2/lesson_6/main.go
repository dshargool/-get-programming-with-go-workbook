package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	var pi32 float32
	fmt.Println(pi32)
	var pi64 float64 = math.Pi
	fmt.Println(pi64)

	// 4 is the width.  5 is the precision when printing
	fmt.Printf("%8.5f\n", pi64)

	// Let's look at how floating point rounding errors show up
	piggyBank := 0.0
	for i := 0; i < 11; i++ {
		piggyBank += 0.1
	}
	// Print 1.0999999999...
	fmt.Println(piggyBank)

	// piggy.go
	total := 0.0
	desired := 20.0
	for total <= desired {
		coin := rand.Intn(3)
		newval := 0.00
		switch coin {
		case 0:
			newval = 0.05
		case 1:
			newval = 0.10
		case 2:
			newval = 0.25
		}
		total += newval
		fmt.Printf("$%7.2f\n", total)
	}
}

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const mars_dist = 62100000
	const seconds_to_days = 60 * 60 * 24

	fmt.Printf("%-20v %-10v %-13v %4v\n", "Spaceline", "Days", "Trip type", "Price")
	spaceline := "None"

	for tickets := 0; tickets < 10; tickets++ {
		num := rand.Intn(3)
		switch num {
		case 0:
			spaceline = "Space Adventures"
		case 1:
			spaceline = "SpaceX"
		case 2:
			spaceline = "Virgin Galactic"
		}
		ship_class := rand.Intn(15) + 1
		fmt.Println(ship_class)
		ship_speed := ship_class + 13
		ship_cost := ship_class + 35
		duration_days := mars_dist / ship_speed / seconds_to_days
		round_trip := rand.Intn(2)
		round_str := "One-way"
		if round_trip == 1 {
			round_str = "Round-trip"
			duration_days *= 2
			ship_cost *= 2
		}
		fmt.Printf("%-20v %-10v %-13v %4v \n", spaceline, duration_days, round_str, ship_cost)
	}
}

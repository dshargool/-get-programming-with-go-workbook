package main

import (
	"fmt"
	"math"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat, long float64
}

type world struct {
	radius float64
}

//Convert from d/m/s to decimal degrees
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func main() {
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}

	fmt.Println(lat, long)
	fmt.Println(lat.decimal(), long.decimal())

	// Composite Literal
	curious := location{lat.decimal(), long.decimal()}
	// Constructor function
	curious_new := newLocation(lat, long)

	fmt.Println(curious)
	fmt.Println(curious_new)

	// Listing 22.4
	var mars = world{radius: 3389.5}

	spirit := location{-14.5684, 175.472636}
	opportunity := location{-1.9462, 354.4734}
	curiosity := newLocation(coordinate{4, 35, 22.2, 'S'}, coordinate{137, 26, 30.1, 'E'})
	insight := newLocation(coordinate{4, 30, 0.0, 'N'}, coordinate{135, 54, 0.0, 'E'})

	var landings = []location{spirit, opportunity, curiosity, insight}

	var low_dist float64 = math.MaxFloat64
	var high_dist float64
	var close_1, close_2, far_1, far_2 int

	for i, landing_1 := range landings {
		for j, landing_2 := range landings {
			if i != j {
				dist := mars.distance(landing_1, landing_2)
				if low_dist > dist {
					low_dist = mars.distance(landing_1, landing_2)
					close_1 = i
					close_2 = j
				}
				if high_dist < dist {
					high_dist = mars.distance(landing_1, landing_2)
					far_1 = i
					far_2 = j
				}

			}
		}
	}
	fmt.Println("Closest are: ", landings[close_1], " and ", landings[close_2], " with distance ", low_dist)
	fmt.Println("Furthest are: ", landings[far_1], " and ", landings[far_2], " with distance ", high_dist)

	dist := mars.distance(spirit, opportunity)
	fmt.Printf("%.2f km\n", dist)
}

package main

import (
	"fmt"
	"math"
)

// By not specifying the field we automatically forward the fields and methods of the underlying types to the struct
type report struct {
	sol         //int
	temperature //temperature
	location    //location
}

type temperature struct {
	high, low celsius
}

func (t temperature) average() celsius {
	return (t.high + t.low) / 2
}

/*func (r report) average() celsius {
	return r.temperature.average()
}*/

func (s sol) days(s2 sol) int {
	days := int(s2 - s)
	if days < 0 {
		days = -days
	}
	return days
}

type location struct {
	lat, long float64
	name      string
}

func (l location) description() string {
	return fmt.Sprintf("Name %v, Lat %v, Long %v", l.name, l.lat, l.long)
}

type celsius float64
type sol int

type gps struct {
	start       location
	destination location
	world
}

func (g gps) message() string {
	str := fmt.Sprintf("%v km remaining", g.distance(g.start, g.destination))
	return str
}

type world struct {
	radius float64
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

type rover struct {
	gps
	name string
}

func main() {
	bradbury := location{-4.5895, 137.4417, "bradbury"}
	t := temperature{high: -1.0, low: -78.0}
	report := report{sol: 15, temperature: t, location: bradbury}

	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %+v C\n", report.temperature.high)
	fmt.Printf("average temp %+v C\n", report.temperature.average())
	fmt.Printf("a balmy %+v C\n", report.high)
	fmt.Printf("average temp %+v C\n", report.average())

	fmt.Println(report.sol.days(1446))
	fmt.Println(report.days(1446))

	// gps.go
	mars := world{3389.5}
	start := location{-4.5895, 137.4417, "bradbury"}
	destination := location{4.5, 135.9, "elysium"}
	gps := gps{start, destination, mars}
	fmt.Println(gps.message())
	fmt.Println(gps.start.description())
	fmt.Println(gps.destination.description())

	curiousity := rover{gps, "curiousity"}
	fmt.Println(curiousity.message())

}

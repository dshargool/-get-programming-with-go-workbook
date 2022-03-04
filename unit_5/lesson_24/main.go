package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type martian struct{}

func (m martian) talk() string {
	return "nick nick"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("new ", int(l))
}

type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type starship struct {
	laser
}

type rover string

func (r rover) talk() string {
	return string(r)
}

// marshall.go

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) String() string {
	return fmt.Sprintf("%vdeg%v'%.1f\" %c", c.d, c.m, c.s, c.h)
}

type location struct {
	Lat, Long coordinate
}

func (l location) String() string {
	return fmt.Sprintf("%v %v", l.Lat, l.Long)
}

func (c coordinate) Decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		DD  float64 `json:"decimal"`
		DMS string  `json:"dms"`
		D   float64 `json:"degrees"`
		M   float64 `json:"minutes"`
		S   float64 `json:"seconds"`
		H   string  `json:"hemisphere"`
	}{
		DD:  c.Decimal(),
		DMS: c.String(),
		D:   c.d,
		M:   c.m,
		S:   c.s,
		H:   string(c.h),
	})
}

func main() {
	var t interface {
		talk() string
	}

	t = martian{}
	fmt.Println(t.talk())

	t = laser(3)
	fmt.Println(t.talk())
	shout(martian{})
	shout(laser(2))
	s := starship{laser(3)}
	fmt.Println(s.talk())
	shout(s)

	r := rover("whir whir")
	shout(r)

	// Marshal.go

	elysium := location{
		Lat:  coordinate{4, 30, 0.0, 'N'},
		Long: coordinate{135, 54, 0.0, 'E'},
	}
	fmt.Println(elysium)
	byte, _ := json.MarshalIndent(elysium, "", "    ")
	fmt.Println(string(byte))

}

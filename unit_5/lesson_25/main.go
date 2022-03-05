package main

import (
	"fmt"
	"math/rand"
)

type animal struct {
	name string
}

type elephant struct {
	animal
}

func (e elephant) String() string {
	return e.name
}

func (e elephant) move() string {
	return "stomps over"
}

func (e elephant) eat() string {
	return "monches grass"
}

type snake struct {
	animal
}

func (s snake) move() string {
	return "slithers"
}

func (s snake) String() string {
	return s.name
}

func (s snake) eat() string {
	return "swallows mouse"
}

type animaler interface {
	String() string
	move() string
	eat() string
}

func step(a animaler) string {
	r := rand.Intn(2)
	var action string
	switch r {
	case 0:
		action = a.eat()
	case 1:
		action = a.move()
	}
	return action
}

func main() {
	ark := []animaler{elephant{animal{name: "Francis"}},
		snake{animal{name: "Sam"}},
	}

	const sunrise = 8
	const sunset = 20

	for i := 0; i < 72; i++ {
		if i%24 < sunrise || i%24 > sunset {
			fmt.Println("Sleep time")
		} else {
			j := rand.Intn(len(ark))
			fmt.Printf("Time: %+v \t Action: %v %v\n", i%24, ark[j].String(), step(ark[j]))
		}
	}
}

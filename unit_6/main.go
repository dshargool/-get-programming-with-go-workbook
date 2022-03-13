package main

import (
	"fmt"
	"sort"
)

func sortStrings(s []string, less func(i, j int) bool) {
	if less == nil {
		less = func(i, j int) bool { return s[i] < s[j] }
	}
	sort.Slice(s, less)
}

func mirepoix(ingredients []string) []string {
	return append(ingredients, "onion", "carrot", "celery")
}

type character struct {
	name     string
	leftHand *item
}

type item struct {
	name string
}

func (c *character) pickup(i *item) string {
	if c.leftHand == nil {
		c.leftHand = i
	} else {
		return fmt.Sprintf("%#v's hands are full", c.name)
	}
	return fmt.Sprintf("%v picks up %v\n", c.name, c.leftHand.name)
}

func (c *character) giveto(i *item, oc *character) {
	if c.leftHand == i {
		c.leftHand = nil
	} else {
		fmt.Printf("%v does not have %v\n", c.name, i.name)
	}

	if oc.leftHand == nil {
		oc.leftHand = i
	} else {
		fmt.Printf("%v has no room to hold %v and dropped it.\n", oc.name, i.name)
	}
}

func main() {
	var nowhere *int
	fmt.Println(nowhere)
	if nowhere != nil {
		fmt.Println(*nowhere) // Causes a panic
	}

	food := []string{"onion", "carrot", "celery"}
	sortStrings(food, nil) // Sort with the default
	fmt.Println(food)
	sortStrings(food, func(i, j int) bool { return len(food[i]) < len(food[j]) }) // Sort with the default
	fmt.Println(food)

	soup := mirepoix(nil)
	fmt.Println(soup)

	// knights.go
	arthur := character{name: "Arthur"}
	knight := character{name: "Robinson"}
	sword := item{"Excalibur"}

	fmt.Printf(arthur.pickup(&sword))
	fmt.Println(arthur)
	fmt.Println(knight)
	arthur.giveto(arthur.leftHand, &knight)
	fmt.Println(arthur)
	fmt.Println(knight)

}

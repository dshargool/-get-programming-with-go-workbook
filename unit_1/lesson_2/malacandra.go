package main

import "fmt"

func main() {
	var distance = 56000000
	var days = 28
	var hours = days * 24

	var speed = distance / hours
	fmt.Println(speed)
}

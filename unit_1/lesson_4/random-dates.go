package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for count := 0; count < 10; count++ {
		era := "AD"
		year := rand.Intn(2022) + 1
		month := rand.Intn(11) + 1
		daysInMonth := 31
		switch month {
		case 2:
			if year%400 == 0 || year%4 == 0 && year%100 != 0 {
				daysInMonth = 29
			}
			daysInMonth = 28
		case 4, 6, 9, 11:
			daysInMonth = 30
		}
		day := rand.Intn(daysInMonth) + 1
		fmt.Println(era, year, month, day)
	}
}

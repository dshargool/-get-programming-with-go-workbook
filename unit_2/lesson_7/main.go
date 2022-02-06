package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// Types of ints: int and uint
	// int8, int16, int32, int64
	// uint8, uint16, uint32, uint64
	//
	// Regular int and uint will be architecture specific

	var myInt int8 = 5
	// Printf %T will show variables type.
	fmt.Printf("%T\n", myInt)

	// CSS uses 8 bits for colour so a good use for representing these colours is uint8
	// This can lead to significant memory savings over the default 32 or 64 bits.

	// Hexadecimal digit consists of 4 bits (a nibble).
	// To use hex in Go you can prefix with 0x
	var myHex uint8 = 0xFF
	fmt.Printf("%x\n", myHex)

	// Integer wrap around
	// Signed ints will overflow.  127 <+-> -128
	// Example of overflow
	myInt = 127
	fmt.Println(myInt)
	myInt++
	fmt.Println(myInt)

	// piggy.go experiment
	// Piggy Bank with cents tracked as an integer instead of dollars

	var coin int
	current_dollar := 0
	current_cent := 0
	desired_dollar := 20

	for current_dollar < desired_dollar {
		coin = rand.Intn(3)
		switch coin {
		case 0:
			coin = 5
		case 1:
			coin = 10
		case 2:
			coin = 25
		}
		current_cent += coin
		if current_cent >= 100 {
			current_dollar++
			current_cent -= 100
		}
		fmt.Printf("$%d.%02d\n", current_dollar, current_cent)
	}

}

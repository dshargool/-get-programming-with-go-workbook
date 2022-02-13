package main

import "fmt"

func main() {
	// Raw string literal
	fmt.Println("Here is a\n line break")
	fmt.Println(`Here is a\n raw string
	literal`)

	// To print the characters rather than the numeric values use the %c format verb with Printf
	// This prints the pi character and A
	fmt.Printf("%c%c\n", 960, 65)

	// Individually access characters
	message := "guinea pig"
	for i := 0; i < len(message); i++ {
		c := message[i]
		fmt.Printf("%c\n", c)
	}

	//ROT13
	message = "uv vagreangvbany fcnpr fgngvba"
	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}

	//caesar.go
	// Decipher the message for caeser
	message = "L fdph, L vdz, L frqtxhuhg"
	fmt.Println()
	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c - 3
			if c < 'a' {
				c = c + 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c = c - 3
			if c < 'A' {
				c = c + 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()

	// international.go
	// ROT13 for spanish
	message = "Hola Estación Espacial Internacional"
	for _, c := range message {
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c = c + 13
			if c > 'Z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
		fmt.Printf("%c", c)
		fmt.Printf("%c", c)
	}
}

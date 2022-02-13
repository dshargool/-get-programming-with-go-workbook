package main

import (
	"fmt"
	"strings"
)

func main() {
	message := "CSOITEUIWUIZNSROCNKFD"
	message = strings.ToUpper(message)
	decrypt := "GOLANG"
	decrypt = strings.ToUpper(decrypt)

	for i := 0; i < len(message); i++ {
		c := message[i]
		dec := byte(decrypt[i%len(decrypt)] - 'A')
		if c >= 'a' && c <= 'z' {
			c = c - dec
			if c > 'z' {
				c = c - 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c = c - dec
			if c > 'Z' {
				c = c - 26
			} else if c < 'A' {
				c = c + 26
			}

		}
		fmt.Printf("%c", c)
	}
}

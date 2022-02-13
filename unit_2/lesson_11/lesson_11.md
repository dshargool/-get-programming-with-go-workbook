# Capstone: The Vigenere Cipher
## Caesar Cipher
Caesar Cipher is known from earlier section.  It shifts the letter by 3 (ahead to encrypt, back to decrypt).  

The message we're trying to decode is:
`LFDPHLVDZLFRQTXHUHG`
which when run through our program gives us:
`ICAMEISAWICONQUERED`


Program:
```go
func main() {
	message := "LFDPHLVDZLFRQTXHUHG"
	decrypt := "D"
	for i := 0; i < len(message); i++ {
		c := message[i]
		// This only works for upper case...
		decryp := byte(decrypt[i%len(decrypt)] - 'A')
		if c >= 'a' && c <= 'z' {
			c = c - decryp
			if c > 'z' {
				c = c - 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c = c - decryp
			if c > 'Z' {
				c = c - 26
			}

		}
		fmt.Printf("%c", c)
	}
}

```

## Vigenere Cipher
Similar to the Caesar Cipher except for a constant value of 3 we use a keyword and apply it across the message and use it's 'value' to decrypt that character.
```
CSOITEUIWUIZNSROCNKFD
GOLANGGOLANGGOLANGGOL
```
Which outputs:
`WEDIGYOULUVTHEGOPHERS`

Program:
```go 
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
```

# Multilingual Text
- Access and manipulate individual letters
- Cipher and decipher secret messages
- Write programs for Multilingual world

## Strings
``` go
peace := "peace"
var peace = "peace"
var peace string = "peace"
```
The zero value for a string is "".

### Raw string literals
Backticks "\`" indicate a raw string literal.

## Characters, code points, runes, bytes
Go provides the type `rune` which is an alias for `int32` to represent 'code points'.  'Code points' are the integer that represents a specific character.  For example the code point 65 represents A.

Go also provides `byte` which is an alias for `uint8` and can be used for English characters within ASCII (128 character subset of Unicode)

#### (Aside) Assigning own aliases
As of Go v1.9 can define own types:
```go
type byte = uint8
type rune = int32
```

## Character literal
To assign a character literal just use single quites.  If no type is specified Go will infer a rune. Making the following equivalent:
```go
grade := 'A'
var grade = 'A'
var grade rune = 'A'
```
Character literals can also be used with the `byte` alias assuming they are within ASCII (not sure what happens if they're not)

## Immutability
Strings are immutable in Go (Same as in Python, Java, etc.) so we can access and change specific characters (like we can in C char arrays).

## UTF-8
UTF-8 is variable length encoding meaning in encapsulates ASCII's 128 bits super well.

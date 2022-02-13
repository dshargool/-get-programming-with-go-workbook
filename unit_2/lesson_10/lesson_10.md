# Lesson 10
Type Conversions

## Types don't mix
Can't combine types.  Numbers add.  Strings concetenate.
```go
var := "This is" + "ok"
var := "This is" + 10 + "not ok" // Compiler error!
```

### Numeric Type Conversions
Convert by casting value
Syntax:
```go
age := 41 // This is an int
floatAge := float64(age) // This casts is to a float64
```

#### Float to Integer
Decimals get truncated without rounding
36.43 => 36 
36.99 => 36

#### Changing sizes
Changing from smaller size `uint8` to larger `uint32` is ok.  Changing to smaller isn't as there might be overflow.  Similar to how going from `int` to `uint` if the integer is negative.  This is why type casting is explicit so that consequences can be considered.

#### Checking Conversions
Can check conversions using math package.  MinInt and MaxInt are unsigned so can be included in comparisons!
```go 
if myVar < math.MinInt16 || bh > math.MaxInt16 {
	// Out of range
}
```

### String Conversions
Can convert rune or byte to string and gives result using `%c` format verb.

#### ASCII Conversions
```go 
strconv.Itoa(intToConvert)
```
```go 
strconv.Atoi(asciiToConvert)
```

### Static types
Types for variables are static so we can't change a variable type from int to float
```go 
var countdown = 10 //int
countdown = 0.5 //float - Error because countdown is typed as in on declaration
fmt.Println(countdown)
```
### Boolean Conversions
Can convert boolean values the same way as other variables. If we want nice strings a simple if statement is suggested
```go 
	launch := false
	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("Ready for launch:", launchText)
	var yesNo string
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	fmt.Println("Ready for launch:", yesNo)
```

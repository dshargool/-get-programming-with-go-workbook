package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	// Playing with types
	myVar := "This is" + "10" + "ok"
	fmt.Println(myVar)
	// myVar = "This is" + 10 + "not ok"

	var bh float64 = 32769
	var h = int16(bh)
	fmt.Println(h)

	// Can use math package to check variable validity before changing.
	// MinInt and MaxInt are untyped so can be easily compared.
	if bh < math.MinInt16 || bh > math.MaxInt16 {
		fmt.Println("Variable out of range")
	}

	// Quick check 10.3: Determine if variable v is within range of uint8

	v := 1230

	if v >= 0 && v <= math.MaxUint8 {
		fmt.Println("Variable in Uint8 range")
	} else {
		fmt.Println("Variable not in Uint8 range")
	}

	///
	///
	///
	// String Conversions

	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33
	fmt.Print(string(pi), string(alpha), string(omega), string(bang), "\n")

	// Integer to ascii
	countdown := 10
	str := "Launch in T Minus " + strconv.Itoa(countdown) + " Seconds."
	fmt.Println(str)

	// Can also use sprintf which returns a string instead of printing it
	countdown = 9
	str = fmt.Sprintf("Launch in T minus %v seconds.", countdown)
	fmt.Println(str)

	// Can also convert ascii to int but may return an error
	countdown, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println("This is bad")
	}
	fmt.Println(countdown)

	// Boolean Conversion
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

	// Summary Experiment
	myString := "ye"
	var bool_val bool
	switch myString {
	case "true", "yes", "1":
		bool_val = true
	case "false", "no", "0":
		bool_val = false
	default:
		fmt.Println("Invalid value")
	}
	fmt.Println(bool_val)

}

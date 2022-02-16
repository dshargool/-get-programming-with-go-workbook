package main

import "fmt"

type celsius float64
type fahrenheit float64
type kelvin float64

func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}
func celsiustoKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

// Can also associate methods with declared type.
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
func (k kelvin) fahrenheit() fahrenheit {
	return fahrenheit(k.celsius().fahrenheit())
}
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((9.0 / 5.0 * c) + 32.0)
}
func (f fahrenheit) kelvin() kelvin {
	return kelvin(f.celsius() + 273.15)
}
func (f fahrenheit) celsius() celsius {
	return celsius((f - 32.0) * (5.0 / 9.0))
}

func main() {
	var temperature celsius = 20.0
	fmt.Println(temperature)

	var zero_kelvin kelvin = 0
	temperature = kelvinToCelsius(zero_kelvin)
	fmt.Println(temperature)

	var moon_temperature celsius = 127
	fmt.Println(celsiustoKelvin(moon_temperature))
	fmt.Println(moon_temperature.kelvin())

	// Test all of our type methods
	fmt.Println("In C: ", temperature)
	fmt.Println("In F: ", temperature.fahrenheit())
	fmt.Println("In K: ", temperature.kelvin())

	temperature = 20
	fmt.Println("In C: ", temperature)
	fmt.Println("In F: ", temperature.fahrenheit())
	fmt.Println("In K: ", temperature.kelvin())
}

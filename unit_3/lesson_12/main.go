package main

import "fmt"

func kelvinToCelsius(k float64) float64 {
	celsius := k - 273.15
	return celsius
}

func celsiusToFahrenheit(c float64) float64 {
	fahrenheit := (c * 9.0 / 5.0) + 32.0
	return fahrenheit
}

func kelvinToFahrenheit(k float64) float64 {
	ret := celsiusToFahrenheit(kelvinToCelsius(k))
	return ret
}

func main() {
	kelvin := 233.0
	celsius := kelvinToCelsius(kelvin)
	zero_kel := 0.0
	fmt.Println(kelvinToFahrenheit(zero_kel))
	fmt.Println(celsius)
}

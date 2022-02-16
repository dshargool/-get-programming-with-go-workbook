package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	//TODO: Implement real sensor
	return 0
}

func measureTemperature(samples int, sensor func() kelvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Printf("%vK\n", k)
		time.Sleep(time.Second)
	}
}

// Type functions
type sensor func() kelvin

// Playing with anonymous functions
func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	sensor := fakeSensor
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())

	// Can also declare as
	var sensor_2 func() kelvin

	sensor_2 = fakeSensor
	my_int := kelvin(5)

	measureTemperature(5, sensor_2)

	sensor = calibrate(realSensor, my_int)
	fake_sen := calibrate(fakeSensor, my_int)
	sensor = func() kelvin {
		return my_int
	}
	fmt.Println(sensor())
	my_int++
	fmt.Println(sensor())
	fmt.Println(fake_sen())

}

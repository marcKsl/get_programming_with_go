package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	var offset kelvin = 10
	sensor := calibrate(realSensor, offset)
	sensor2 := fakeSensor
	fmt.Println(sensor2())
	fmt.Println(sensor2())
	fmt.Println(sensor2())
	fmt.Println(sensor2())
	fmt.Println(calibrate(sensor2, offset)())
	fmt.Println(sensor2())
	fmt.Println(sensor())
	// Changing the value of offset won't affect the output of the sensor() function
	// The initial offset value was passed in by value, so it does not reference the offset value
	// anymore.
	offset = 20
	fmt.Println(sensor())
}

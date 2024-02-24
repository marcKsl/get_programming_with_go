package main

import (
	"fmt"
)

type kelvin float64

// sensor function type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	// We return a anonymous function, or respectively a function literal
	return func() kelvin {
		// We can use s and offset of a surrounding scope because
		// anonymous functions enclose the variables in scope
		// These variables are accessible to the anonymous function
		// even after the calibrate function has ended.
		return s() + offset
	}
}

func main() {
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor())
	var k kelvin = 297.0
	// The k inside the anonymous function references the k above
	tempKelvin := func() kelvin { return k }
	fmt.Println(tempKelvin())
	// Therefore it can be incremented
	k++
	// And the function uses the incremented value, because it's like said above, just a reference
	fmt.Println(tempKelvin())
	// This seems to come in handy in for loops
}

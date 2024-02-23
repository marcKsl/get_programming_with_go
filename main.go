package main

import (
	"fmt"
)

type celsius float64
type kelvin float64

// Typical function declaration
func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

// Typical method declaration
// Note: Methods can, like functions, accept multiple parameters, but must have exactly one receiver
// The receiver is the variable the method gets called on.
// It needs to be of the type, specified in the receiver part of the method.
// The body is the same.
// Unlike in functions, in methods both the receiver and parameters stand in front of the method name.
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func main() {

	var c celsius = 127 // sunlit moon surface temperature
	var k kelvin

	// Typical function call
	k = celsiusToKelvin(c)
	fmt.Printf("%v\n", k)
	c = 20 // room temperature on earth
	// Typical method call.
	// Methods are called with the dot notation on variables of the correct type,
	// followed by the method name and any arguments.
	k = c.kelvin()
	fmt.Printf("%v\n", k)
}

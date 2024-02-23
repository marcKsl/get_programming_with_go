package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0
}

func main() {
	fmt.Println("Hello World")
	// Functions in go a first-class, they work everywhere where integers, strings, and other types work.
	// Functions can be assigned to variables.
	// Note: The function itself and not its return is assigned to the variable.
	// The sensor variable is now of type "function", that accepts no parameters and has a return type of kelvin.
	// Note: The last part is relevant, you can not assign different function types after that to the same variable.
	// If we would not rely on type inference, we need to write the declaration like this:
	// var sensor func() kelvin
	sensor := fakeSensor
	fmt.Println(sensor())
	// This way we have isloated the actual sensor and can switch it without affecting the code at all.
	// Pretty, pretty neat, i must say.
	// Note: As mentioned above, we can not assign a function that accepts also no parameters but returns the type celsius for example.
	sensor = realSensor
	fmt.Println(sensor())
	// Functions can be passed to other functions as arguments.
	// You can write functions that create other functions.

}

package main

import (
	"fmt"
)

type celsius float64
type fahrenheit float64

func celciusToFahrenheit(c celsius) fahrenheit {
	return fahrenheit((c * 9.0 / 5.0) + 32.0)
}

func main() {
	var temperature celsius = 20
	// Values like 21.4 are also considered constant values.
	// Constants can be assigned to the newly created Celsius
	// type because they are untyped and infer their type
	// from the context in which they are used.
	// This behavior is similar to that of float64.
	const degrees = 20
	temperature++
	temperature += 21.4
	temperature += degrees
	// The newly created Celsius type is unique though.
	// It's not a type alias for float64,
	// so in the following statement,
	// the float64 value first needs to
	// be converted into the Celcius type.
	var myFloaty float64 = 20.4
	// As explained, this cannot be used directly.
	// It first needs to be converted.
	temperature += celsius(myFloaty)
	fmt.Println(temperature)
	// Creating types can be useful for improving readability and reliability.
	// Readability, because you can literally read the type of the variable and know what it does.
	// In terms of reliability, you can detect unwanted operations between different temperature types, for example.
	var temperatureInFahrenheit fahrenheit = 50
	fmt.Println(temperatureInFahrenheit)
	// We first need to convert the type to add both of the values together.
	// So you can for example use a function to convert these
	// types and therefore ensure correct behaviour during the conversion.
	higherTemperature := celciusToFahrenheit(temperature) + temperatureInFahrenheit
	fmt.Println(higherTemperature)
	// This will print main.fahrenheit because the types declared inside a package are scoped to it.
	// The package name is an essential part of the fully qualified name of a type.
	// Only primitives are built-in types and do not belong to a specific package.
	fmt.Printf("%T", higherTemperature)
}

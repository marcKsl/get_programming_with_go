package main

import (
	"fmt"
)

// kelvinToCelcius converts Kelvin to Celcius
func kelvinToCelcius(k float64) float64 {
	return k - 273.15
}

func main() {
	fmt.Println("Hello World")

	// Note: functions from the same package are invoked without prefixing the package name
	// Note: We do not "give" the original variable to the function,
	// instead we give the value of the variable k in the main scope to the variable k in the function scope
	// we could also name the variables kScope and kFunction, it would always work. This is called pass by value.
	k := 350.44
	c := kelvinToCelcius(k)
	fmt.Println(c)

	// A good reason to use functions is
	// that they allow you to encapsulate a
	// piece of logic in a safe manner and
	// reuse it throughout your code with a meaningful name.

	// Summary
	//-----------------------------------
	// Functions are declared with a name, parameters and a list of returns
	// Capitalized function names, types and variables are made available to other packages
	// Each parameter and result is followed by a type, types can be chained like: func MyFunction(n, m int) (e, n int)
	// If you want to call a function in another package than the package it was declared in,
	// you need to prefix it with the corresponding package name
	// Functions are called with arguments that correspong to the parameters they accept.
	// Results are returned to the caller with the return keyword

}

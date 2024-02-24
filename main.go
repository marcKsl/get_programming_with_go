package main

import "fmt"

func main() {
	// Arrays are of a fixed size in Go.
	// If you try to access an array out of bounds,
	// the Go compiler will report an error.
	var planets [8]string
	// These lines do not work in Go:
	// planets[8] = "Pluto"
	// pluto := planets[8]
	// Be careful, the above lines do not work in Go.
	// If for some reason the go compiler cannot detect the error,
	// the program will panic during runtime. That is good behaviour
	// because it ensures that no unspecified memory operation will be performed,
	// as is the case in c
	// Always stay safe and access arrays inside their boundaries:
	planets[7] = "Neptune"
	fmt.Println(planets[7])
}

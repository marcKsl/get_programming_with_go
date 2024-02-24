package main

import (
	"fmt"
)

func main() {
	// This declares a string array named planets with a length of 12
	var planets [12]string
	// The following lines assign different values to different indices in the given array.
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth"
	planets[3] = "Mars"
	planets[4] = "Jupiter"
	planets[5] = "Saturn"
	planets[6] = "Uranus"
	planets[7] = "Neptune"

	// We can retrieve the value at a given index in the array like so:
	earth := planets[2]
	// Or we can use it directly in function calls as arguments:
	fmt.Printf("The people from %v would love to live on %v one day.\n", earth, planets[3])

	// We can retrieve the length of an array like so:
	fmt.Println(len(planets))
	// The len() function is a built-in function in go and does not require any package imports.

	// Elements of an Array are initialized with the zero value of the array's type, for example 0 for integer arrays.
	// A complete list of all zero values:
	// int, int8, int16, int32, int64, uint, uint8 (alias byte), uint16, uint32, uint64, uintptr: 0
	// float32, float64: 0.0 <- Note that it includes one floating point
	// complex64, complex128: 0+0i
	// bool: false
	// string: "" (leerer String) <- That's kind of interesting but i like it a lot
	// pointers, functions, interfaces, slices, channels, maps: nil <- i really like the word nil, to be honest.
}

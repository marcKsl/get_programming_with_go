package main

import (
	"fmt"
)

// Read on, come back here later when told so.
// This function is completely useless.
// It will take in an array as a parameter.
// By nature in go, functions use pass by value, which means
// the original array will not be referenced,
// instead we modify a newly created array, that holds the same values
// as the array that was given in as an argument.
// But this newly created array is no longer in memory when the function ends,
// and the original array has never been modified.
func terraform(planets [8]string) {
	for i := range planets {
		planets[i] = "New" + planets[i]
	}
}

func main() {
	fmt.Println("Hello World")

	planets := [8]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	moonAndEarth := [2]string{"Moon", "Earth"}

	// Assigning an array to a new variable will make a copy of it.
	planets2 := planets

	planets[2] = "Azeroth"

	fmt.Println(planets)
	// No Azeroth here, still the good ol' earthy
	// Because we do not reference the original array,
	// instead we reference a newly created array,
	// that contains the same values as planets.
	fmt.Println(planets2)

	// This does not change / do anything, see explanation above.
	terraform(planets)
	fmt.Printf("No new planets at all: %v\n", planets)

	// It is also important to notice, that the length of an array is part of its type.
	// So if we have two arrays, both are collections of strings,
	// but don't share the same length, we cannot assign one to each other.
	fmt.Println(moonAndEarth)
	fmt.Printf("The moonAndEarth array is: %v long.\n", len(moonAndEarth))
	fmt.Printf("The planets array is: %v long.\n", len(planets))
	fmt.Println("")
}

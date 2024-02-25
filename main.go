package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	// This is one way of delcaring and intializing a slice.
	// Take an array and slice the whole thing afterwards.
	dwarfArray := [...]string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfSlice := dwarfArray[:]
	fmt.Println(dwarfSlice)

	// We can also declare a slice directly.
	// A slice of strings has the type []string, unlike arrays,
	// which always have a number or ellipsis between the brackets.
	// This is due to the fact, that arrays are of a fixed length, and slices not.
	// Slices reference an array, specify their length and capacity.
	// The capacity is similiar to the length of the referenced array.
	// If the slice grows beyond the length of the array, a new array is created by go,
	// and the slice references the new one from now on.

	// To delcare a slice directly we can use the composite literal again, as we did with arrays.
	// Note: The only visible difference are the empty brackets.
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Println(dwarfs)

	fmt.Printf("dwarfArray is of type: %T\ndwarfs is of type: %T", dwarfArray, dwarfs)
}

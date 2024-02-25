package main

import "fmt"

func main() {
	planets := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}
	// This will create a slice of integers.
	// Note: It will go up to the 4th index, but it will not include it.
	terrestrial := planets[0:4]
	gasGiants := planets[4:6]
	iceGiants := planets[6:8]

	fmt.Println(planets)
	fmt.Println(terrestrial)
	fmt.Println(gasGiants)
	fmt.Println(iceGiants)

	// We can normally index into slices, like with arrays:
	fmt.Println(terrestrial[2])

	// Slices can be sliced into more slices:
	giants := planets[4:8]
	gas := giants[0:2]
	ice := giants[2:4]

	fmt.Println(giants)
	fmt.Println(gas)
	fmt.Println(ice)

	// A slice always references an array,
	// so if we modify the slice, we modify the underlying array,
	// so that the change will be visible in all other slices too.
	giants[0] = "Oops, i di..."
	fmt.Println(planets)
	fmt.Println(giants)
	fmt.Println(gas)
	fmt.Println(ice)

	// If we omit the first index when slicing an array, it defaults to the beginning of the array.
	// If we omit the last index of the array, it defaults to the length of the array.

	terrestrial2 := planets[:4]
	gasGiants2 := planets[4:6]
	iceGiants2 := planets[6:]
	fmt.Println(terrestrial2)
	fmt.Println(gasGiants2)
	fmt.Println(iceGiants2)

	// Ommiting both indices leads to a slice containing the whole array.
	allPlanets := planets[:]
	fmt.Println(allPlanets)

	// Slice indices may not be negative.

	// Strings can also be sliced.

	neptune := "Neptune"
	// This unlike with arrays creates a new variable and no reference.
	tune := neptune[3:]
	fmt.Println(tune)
	// So this new assertion will not affect the "tune" variable.
	neptune = "Moon"
	fmt.Println(tune)

	//  Note: Indices indicate the number of bytes, not runes.
	question := "¿Cómo estás?"
	fmt.Println(question[:6])

	colonized := terrestrial[2:]
	fmt.Println(colonized)
}

package main

import "fmt"

// maps are pointers
// dont do this:
// func demolish(planets *map[string]string)

// slices are also pointers to arrays
// slices almost never need to be adressed as a pointer
// unless you want to mutate the length, the pointer itself or the capacity
// here we change the length of a slice by adressing it directly through its pointer
func reclassify(planets *[]string) {
	*planets = (*planets)[0:8]
}

func main() {
	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
		"Pluto",
	}
	fmt.Println(planets)
	// poor pluto..
	reclassify(&planets)
	fmt.Println(planets)

	// So the lesson is: We do not need pointers to modify the data that maps and slices hold
	// But we do need pointers to modify the data that sits inside of structures and arrays
}

// Lesson 18

package main

import (
	"fmt"
)

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {

	// The append function

	fmt.Println("-------------------------------------------------------")
	fmt.Println("-------------------The append function-----------------")
	fmt.Println("-------------------------------------------------------")

	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	fmt.Println(len(dwarfs))
	dwarfs = append(dwarfs, "Orcus")
	fmt.Println(dwarfs)
	fmt.Println(len(dwarfs))

	// Length and capacity

	fmt.Println("-------------------------------------------------------")
	fmt.Println("-------------------Length and capacity-----------------")
	fmt.Println("-------------------------------------------------------")

	dwarfs = []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dump("dwarfs", dwarfs)
	fmt.Println("")
	dump("dwarfs[1:2]", dwarfs[1:2])

	// Investigating the append function

	fmt.Println("-------------------------------------------------------")
	fmt.Println("------------Investigating the append function----------")
	fmt.Println("-------------------------------------------------------")

	dwarfs = []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dwarfs2 := append(dwarfs, "Orcus")
	dwarfs3 := append(dwarfs, "Salacia", "Quaoar", "Sedna")
	dump("dwarfs", dwarfs)
	dump("dwarfs", dwarfs2)
	dump("dwarfs", dwarfs3)

	// All three slices point to different underlying arrays
	dwarfs3[0] = "Pluto"
	dump("dwarfs", dwarfs)
	dump("dwarfs", dwarfs2)
	dump("dwarfs", dwarfs3)

	// Three index slicing

	fmt.Println("-------------------------------------------------------")
	fmt.Println("------------------Three index slicing------------------")
	fmt.Println("-------------------------------------------------------")

	// a(x:y:z)
	// a is the source slice
	// x is the starting index
	// y is the length or last index
	// z is the capacity

	// mySlice(1:3:5)
	// 1 is the start
	// length is 3-1=2
	// capacity is 5-1=4
	// so it takes the second and third element of the slice and has two more empty slots available

	planets := []string{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	terrestrial := planets[0:4:4]
	worlds := append(terrestrial, "Ceres")

	fmt.Println(planets)
	fmt.Println(terrestrial)
	fmt.Println(worlds)

	terrestrial = planets[0:4]
	worlds = append(terrestrial, "Ceres")
	fmt.Println(terrestrial)
	fmt.Println(worlds)

	// Preallocate slices with make

	fmt.Println("------------------------------------------------------")
	fmt.Println("-------------Preallocate slices with make-------------")
	fmt.Println("------------------------------------------------------")

	dwarfs = make([]string, 0, 10)
	dwarfs = append(dwarfs, "Ceres", "Pluto", "Haumea", "Makemake", "Eris")
	fmt.Println(dwarfs)

	// Declaring variadic functions

	fmt.Println("------------------------------------------------------")
	fmt.Println("-------------Declaring variadic functions-------------")
	fmt.Println("------------------------------------------------------")

	var terraform = func(prefix string, worlds ...string) []string {
		newWorlds := make([]string, len(worlds))

		for i := range worlds {
			newWorlds[i] = prefix + " " + worlds[i]
		}
		return newWorlds
	}

	planets = []string{"Venus", "Mars", "Jupiter"}
	newPlanets := terraform("New", planets...)
	fmt.Println(newPlanets)

	// capacity.go

	fmt.Println("-------------------------------------------------")
	fmt.Println("-------------experiment: capacity.go-------------")
	fmt.Println("-------------------------------------------------")

	var appendAndCount = func(r int, s []string) {
		oldCap := cap(s)
		for i := 0; i <= r; i++ {
			s = append(s, fmt.Sprintf("%v", i))
			newCap := cap(s)
			if oldCap < newCap {
				fmt.Printf("Slice has now a capacity of: %v (Capacity was: %v)\n", newCap, oldCap)
				oldCap = newCap
			}
		}
	}

	startingSlice := make([]string, 0, 10)

	appendAndCount(400, startingSlice)

}

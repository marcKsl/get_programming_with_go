package main

import (
	"fmt"
)

func main() {
	// Summary
	// The Print, Println, Printf functions display text and numbers on the screen
	// The Printf function uses %v format verbs which can be replaced by variables appended to the string
	fmt.Printf("look at this sweet %v ya\n", "variable")

	// Constants are declared with const and can not be changed (immutable)
	// Variables are declared with var and can be assigned new values while a programm is running (mutable)
	var myVarOne, myVarTwo = 12, "12"
	const (
		myConstOne   = true
		myConstTwo   = "onehundred"
		myConstThree = 54.7645643857
	)
	fmt.Println(myVarOne, myVarTwo, myConstOne, myConstTwo, myConstThree)

	// The math/rand import path refers to the rand package
	// The rand.Intn(n) function returns a pseudorandom number between [0,n]
	// Five Go keywords were used in chapter 2: package, import, func, const and var

	// Experiment: malacandra.go

	// Malacandra is much nearer than that: we shall make it in about twenty-eight days.
	// C.S. Lewis, Out of the Silent Planet

	// Malacandra is another name for Mars in The Space Trilogy by C. S. Lewis.
	// Write a program to determine how fast a ship would need to travel (in km/h)
	// in order to reach Malacandra in 28 days. Assume a distance of 56,000,000 km.
	// Compare your solution to the code listing in the appendix.
	const distance, time = 56000000, 28 * 24
	fmt.Println("the velocity needs to be at least:", distance/time, "km/h")
	fmt.Println("or in m/s:", distance/time/3.6)

	// FIN Chapter2

}

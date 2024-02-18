package main

import (
	"fmt"
)

func main() {

	fmt.Println("-------------------------------------------------------------------")
	fmt.Println("---------------------Converting Boolean values---------------------")
	fmt.Println("-------------------------------------------------------------------")

	// Boolean values get converted to their respective text values in strings when printed
	launch := false
	launchText := fmt.Sprintf("%v", launch)
	fmt.Println("Ready for launch:", launchText)

	// To convert them into numbers or other text values, a simple if statement does the trick
	var yesNo string
	if launch {
		yesNo = "yes"
	} else {
		yesNo = "no"
	}
	fmt.Println("Ready for launch:", yesNo)

	// Backwards conversion is more simple, because you can assign the result of a condition
	// directly to a variable.
	// This will print true.
	var myString string = "Yumyum"
	myBooly := (myString == "Yumyum")
	fmt.Println(myBooly)

	// Booleans in go don't have a numeric equivalent,
	// unlike in other languages like C where 0 and 1 are also used.

	// Question: How would you convert a Boolean to an Integer,
	// with 1 for true and 0 for false?

	iWantToBeAnIntegerBool := true
	var iAmNowAnIntegerNotABoolAnyMore int
	if iWantToBeAnIntegerBool {
		iAmNowAnIntegerNotABoolAnyMore = 1
	} else {
		iAmNowAnIntegerNotABoolAnyMore = 0
	}
	fmt.Printf("%v", iAmNowAnIntegerNotABoolAnyMore)

	// Experiment input.go

	fmt.Println("")
	fmt.Println("-------------------------------------------------------------------")
	fmt.Println("------------------------Experiment input.go------------------------")
	fmt.Println("-------------------------------------------------------------------")

	input := "true"
	var output bool
	if input == "1" || input == "yes" || input == "true" {
		output = true
		fmt.Println(output)
	} else if input == "0" || input == "no" || input == "false" {
		output = false
		fmt.Println(output)
	} else {
		fmt.Println("Invalid input.")
	}

}

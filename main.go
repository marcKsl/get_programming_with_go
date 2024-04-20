package main

import "fmt"

func main() {

	answer := 42
	// Referencing the adress of a variable
	fmt.Println(&answer)

	var adress *int

	adress = &answer
	// Dereferencing the adress of a variable to its value
	fmt.Println(*adress)
	// The type of adress
	fmt.Printf("%T\n", adress)

	canada := "Canada"

	var home *string
	fmt.Printf("home is of type: %T\n", home)

	home = &canada
	fmt.Println(*home)
}

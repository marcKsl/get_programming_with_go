// now lets go with strings, or what they are actually called: literal string.

package main

import "fmt"

func main() {

	// all the same, literal values between quotes are inferred to be of the type string.
	peace := "peace"
	var war = "war"
	var defNotWar string = "defNotWar"

	fmt.Println("---------------------------------------------------------------------")
	fmt.Println(peace, war, defNotWar)

	var blank string // Declaration also initializes with the zero value for its type. The zero value for the string type is "".
	fmt.Println(blank)
	fmt.Println("---------------------------------------------------------------------")

	// `` Backticks indicate a raw string literal,e.g. \n wont work in it
	var rawString string = `Escape Character \nis visible. 
	Also these badboys can span multiple lines in go. I bet they aren't used for
	database queries or something like this. Definetly not, no.`
	var normalString string = "Escape Character \nis not visible and breaks line."
	fmt.Println(rawString)
	fmt.Println(normalString)

	// Both, literal and raw strings result in the type string

	fmt.Println("---------------------------------------------------------------------")
	fmt.Printf("normalString is of type: %T\nAnd has the value:\n%[1]v\n", rawString)
	fmt.Println("---------------------------------------------------------------------")
	fmt.Printf("rawString is of type: %T\nAnd has the value:\n%[1]v\n", normalString)
	fmt.Println("---------------------------------------------------------------------")
}

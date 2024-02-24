package main

import (
	"fmt"
)

// Anonymous functions can be assigned to variables.
// These variables can then be used as any other variable.
var f = func() {
	fmt.Println("Hello")
}

// The variable you declare can be in the scope of the package or within a function.

func main() {
	// They can also be assigned inside the main() function like any other function.
	var g = func() {
		fmt.Println("sunny")
	}
	// They can be called like so:
	f()
	g()
	// Of course they can accept parameters
	var h = func(s string) {
		fmt.Println(s)
	}
	// And then get called with arguments
	h("World!")
	// And of course they can specify the return type/s
	var j = func(s string) string {
		return "I " + s + " it!"
	}
	fmt.Println(j("like"))
	var k = func(s string) (string, string) {
		return s, s + "!"
	}
	m1, m2 := k("Hey")
	fmt.Println(m1, m2)

	// We can even declare and call them in one statement.
	func() {
		fmt.Println("One Statement, crazy.")
	}()
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	// Composite literals in go are a special syntax known by the compiler to
	// declare a collection type and directly initialize it with values.
	a := [4]int{0, 1, 2, 3}
	// This line does not compile in go, cause the initialization goes out of bounds
	// c := [3]int{0, 1, 2, 3, 4}
	// This is absolutely fine, the remaining indices get initialized
	// with the zero value of the type, in the case of int it is 0
	b := [7]int{0, 1, 2, 3, 4}
	// We can span the initializationb across multiple lines in go.
	// We can also use the ellipsis to tell go that it
	// should pick the correct size for the array instead of providing an integer
	c := [...]string{
		"mamamia",
		"here i go again",
		"my, my,",
		"how can i resist you?",
	}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(len(c))
}

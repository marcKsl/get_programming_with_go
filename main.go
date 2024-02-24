package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")

	a := [...]string{"Is this the real life?", "Is this just fantasy?"}

	// To iterate over an array there are two possible ways in go.
	// The first one being a normal for loop that uses the len() function
	// to determine the boundaries of the array.
	// Note: It bears the danger to exceed the boundaries of the array if not implemented correctly.
	// Note: The traditional for loop allows for loop customization like accessing only odd elements.
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}
	// The second one involves the range keyword and gives us access to both,
	// the index and the value at the specific index.
	for i, s := range a {
		fmt.Println(i, s)
	}
	// If we do not need the index variable, we can use the blank identifier.
	for _, s := range a {
		fmt.Println(s)
	}
}

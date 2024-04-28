package main

import "fmt"

func main() {

	// This is an anonymous function that will be executed when the function,
	// it sits in returns. Deffered functions will even be executed if the program
	// had a panic.
	// Note: recover() only make sense in deffered functions.
	defer func() {
		e := recover()
		if e != nil {
			fmt.Printf("The following error ocurred: %v", e)
		}
	}()

	fmt.Println("Hello World!")
	// This will give zero the empty type of int, which is indeed 0
	var zero int
	// This will result in a panic
	_ = 42 / zero
	// This won't be executed, because the program had a panic already,
	// and only deferred functions will be executed after a panic.
	a := recover()
	if a != nil {
		fmt.Println(a)
	}
	// This would force a panic, but won't be executed aswell
	panic("PANEK")
}

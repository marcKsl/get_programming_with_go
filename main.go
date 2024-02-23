package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

// A function in go consists of the following parts:
// func FunctioName (n int) int
// The function keyword "func" the function Name,
// the parameters with their respective types and the return type

func MyFunction(n int) int {
	return n + 1
}

func main() {

	// A function call
	// The integer 10 is passed in as an argument to the function call
	fmt.Println(MyFunction(10))

	// This calls the Intn() Function from the rand-Package
	// We can find the definition of functions in the standard package library on: https://pkg.go.dev/std
	fmt.Println(rand.Intn(10))

	// Quick reminder: Functions accept parameters and functions get called with arguments

	// Call of the function Unix(),
	// that returns the local time as type Time,
	// corresponding to the time passed
	// since 01/01/1970 and the given seconds / nanoseconds
	// as arguments to the call
	// Note that both arguments need to be provided, or the
	// compiler throws an error
	fmt.Println(time.Unix(10, 0))

	// A function can be called with multiple parameters like the Atoi function in the strconv package

	s, err := strconv.Atoi("10")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(s)
	}

	// A look into the future
	// The Println() function from the fmt package is a special one
	// It accepts multiple arguments but only one is mandatory
	// And the types are not relevant in the function call
	// The declaration looks like this:
	// func Println(a ...interface{}) (n int, err error)
	// It makes use of two things we don't know yet.
	// One being the "...", which indicates, that this function is
	// variadic, so it can accept a variable number of arguments in its call
	// The next one is the empty interface "interface{}" that lets us every type.
	// Both things will be discussed in further Lessons.

}

package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// This is typical for go
// ErrSomething will be declared and exported at package level
var (
	ErrBounds = errors.New("out of bounds")
	ErrDigit  = errors.New("invalid digit")
)

// We can also create custom error types
type SudokuError []error

// It just needs an Error() string method
func (se SudokuError) Error() string {
	var s []string
	// we cycle through each error in the error slice
	for _, err := range se {
		// basically calling the Error method on every single error
		// and putting it all in a string slice
		s = append(s, err.Error())
	}
	return strings.Join(s, ", ")
}

const rows, columns = 9, 9

// We define the length of the two dimensional array, so that
// we actually create an array and no slice
type Grid [rows][columns]int8

// It is considered good practice to return the error interface
// and not specific types like SudokuError, so that the method
// stays flexible.
func (g *Grid) Set(row, column int, digit int8) error {
	var errs SudokuError
	if !inBounds(row, column) {
		errs = append(errs, ErrBounds)
	}
	if !validDigit(digit) {
		errs = append(errs, ErrDigit)
	}
	if len(errs) > 0 {
		return errs
	}
	g[row][column] = digit
	return nil
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func inBounds(row, column int) bool {
	if row < 0 || row >= rows {
		return false
	}
	if column < 0 || column >= columns {
		return false
	}
	return true
}

func main() {
	var g Grid
	err := g.Set(10, 0, 15)
	if err != nil {
		// This could also be written as:
		// if err, ok := err.(SudokuError); ok {}
		// But i prefer this, because i think it is more concise with
		// the rest of the syntax in go and looks even more readable to me
		// The err.(SudokuError) is a type assertion, that checks if the given
		// Error is of type SudokuError, if so err will get the type SudokuError
		// and ok will be set to true.
		// We need to do this, so that we can correctly handle the error output.
		// Because SudokuError is of type []error, we have to destructure the slice
		// in a for loop to access every actual error plus type and its message.
		err, ok := err.(SudokuError)
		if ok {
			fmt.Printf("%d error(s) occured:\n", len(err))
			for _, e := range err {
				fmt.Printf("- %v", e)
			}
		}
		// Will kill the current process.
		os.Exit(1)
	}
}

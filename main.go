package main

import (
	"fmt"
	"math/rand"
)

var era = "AD" // Declared at package scope, short declaration not possible

func main() {
	year := 2018 // Declaration at function scope, short declaration possible

	switch month := rand.Intn(12) + 1; month { // Declaration of month at switch scope, if switch ends, variable is no longer in scope
	case 2:
		day := rand.Intn(28) + 1 // Each day variable is unique for its own case statement, and is out of scope after the case ends
		fmt.Println(era, year, month, day)
	case 4, 6, 9, 11:
		day := rand.Intn(30) + 1
		fmt.Println(era, year, month, day)
	default:
		day := rand.Intn(31) + 1
		fmt.Println(era, year, month, day)
	}
	mainTwo()
}

// A better approach to reduce code duplication and therefore possible errors

func mainTwo() {
	year := 2018
	month := rand.Intn(12) + 1
	daysInMonth := 31

	switch month {
	case 2:
		daysInMonth = 28
	case 4, 6, 9, 11:
		daysInMonth = 30
	}

	day := rand.Intn(daysInMonth) + 1
	fmt.Println(era, year, month, day)
}

// Summary of Lesson 4:
// New scopes can be built with: {} variables declared in this scope only exist there
// Generally the location where a variable is declared determines its scope
// Shortcut declaration can be used in places where normal declaration can not be used, such as switch and for statements
// Those variables that are declared on the same line as e.g. a for statement, run out of scope when the statement ends
// Both, wide and narrow scopes have their pros and cons
// Narrow scopes can produce duplicate code, but wider scopes can get out of hand pretty quick and become unreadable

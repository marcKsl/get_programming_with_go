package main

import (
	"fmt"
)

func main() {

	// Two ways to implement switch cases, also note the use of fallthrough and default

	// Way number one switch over a variable:

	const floor = 2

	switch floor {
	case 1:
		fmt.Println("This is the first floor.")
	case 2:
		fmt.Println("This is the ground floor.")
		fallthrough
	case 3:
		fmt.Println("This is the basement and the ceiling seems kinda bad")
	default:
		fmt.Println("Still outside!")
	}

	const parallelDimensionFloor = 1

	switch parallelDimensionFloor {
	case 1:
		fmt.Println("This is the first floor.")
	case 2:
		fmt.Println("This is the ground floor.")
		fallthrough
	case 3:
		fmt.Println("This is the basement and the ceiling seems kinda bad")
	default:
		fmt.Println("Still outside?")
	}
}

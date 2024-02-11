package main

import (
	"fmt"
	"math/rand"
	"time"
)

// In Go ONLY Booleans can be used to evaluate a condition
// Go provides branching and repetition with if, switch and for
// So far we have used the following 25 Go keywords:
// package, import, func, var, if, else, switch, case, default, fallthrough, for, break

func main() {
	const myGuess = 89
	var counter = 0

	for {
		time.Sleep(30 * time.Millisecond)
		var cpuGuess = rand.Intn(100) + 1
		if (cpuGuess) == myGuess {
			fmt.Println(myGuess, "is correct!")
			break
		} else if cpuGuess <= myGuess {
			fmt.Println(cpuGuess, "- Nope, your guess is lower than my guess.")
			counter++
		} else {
			fmt.Println(cpuGuess, "- Nope, your guess is higher than my guess.")
			counter++
		}
	}

	fmt.Println("It took you:", counter, "tries to get the right number.")
	if counter < 10 {
		fmt.Println("That was impressive.")
	}

}

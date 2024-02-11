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
	var count = 10

	for count > 0 {
		if rand.Intn(4) == 0 {
			fmt.Println("Something did an oopsie daisy")
			break
		}
		fmt.Println(count)
		time.Sleep(300 * time.Millisecond)
		count--
	}

}

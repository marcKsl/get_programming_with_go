package main

import (
	"fmt"
)

func main() {

	num := 10

	fmt.Println(num)

	// Only possible since GO 1.22
	for i := range 10 {
		num := 33
		fmt.Println(i)
		fmt.Println(num)
	}

	// Before GO 1.22
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		fmt.Println(num)
	}

}

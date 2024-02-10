package main

import (
	"fmt"
)

func main() {

	// 2.4.1

	var a, b = 10, 20

	// 2.4.2

	b += 1
	a /= 5
	b += a
	a++
	a--
	fmt.Println(a, b)

	var weight = 90
	weight -= 2
	fmt.Println("Shortest way to substract 2 from a variable -> 90 -= 2 ->", weight)
}

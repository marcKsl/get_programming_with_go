package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(10) + 1
	fmt.Println(num)

	// Q1
	// The distance between Earth and Mars varies from nearby to opposite sides of the sun.
	// Write a program that generates a random distance from 56,000,000 to 401,000,000 km.
	var variableDistance = rand.Intn(401000001-56000000) + 56000000
	fmt.Println(variableDistance)

}

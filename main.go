package main

import (
	"fmt"
	"math/rand"
)

func getRandomCoin() float64 {
	switch i := rand.Intn(3); i {
	case 0:
		return 0.05
	case 1:
		return 0.10
	default:
		return 0.25
	}
}

func main() {

	piggyBank := 0.0

	for i := 0; piggyBank <= 20; i++ {
		piggyBank += getRandomCoin()
	}

	fmt.Println(piggyBank)
}

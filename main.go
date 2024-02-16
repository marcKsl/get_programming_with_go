package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getCoinValue() int {

	switch i := rand.Intn(3); i {
	case 0:
		return 5
	case 1:
		return 10
	default:
		return 25
	}
}

func main() {
	var piggyBank int = 0
	var goal int = 2000
	fmt.Println(piggyBank)
	for piggyBank <= goal {
		piggyBank += getCoinValue()
		time.Sleep(40 * time.Millisecond)
		fmt.Printf("%.2f$\n", (float64(piggyBank) / 100))
	}
}

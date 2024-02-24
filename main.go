package main

import (
	"fmt"
	"math/rand"
	"time"
)

type kelvin float64
type sensor func() kelvin

func measureTemperature(samples int, s sensor) {
	for i := 0; i < samples; i++ {
		k := s()
		fmt.Printf("%v° K\n", k)
		time.Sleep(time.Second)
	}
}

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

// Question:
// Rewrite the following function signature to use a function type:
// func drawTable(row int, getRow func(row int) (string, string))

// Solution:
// type getRowFn func(row int) (string, string)
// func drawTable(rows int, getRow getRowFn) {
// 	***Implement Function Body here***
// }

func main() {
	measureTemperature(4, fakeSensor)
}

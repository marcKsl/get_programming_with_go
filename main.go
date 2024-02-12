package main

import (
	"fmt"
	"math/rand"
	"time"
)

var era = "AD"

func isLeap(n int) bool {
	return (n%4 == 0 && n%100 != 0) || n%400 == 0
}

func getRandomDate() (int, int, int) {

	year := rand.Intn(2078)
	month := rand.Intn(12) + 1
	daysInMonth := 31

	switch month {
	case 2:
		if isLeap(year) {
			daysInMonth = 29
		} else {

			daysInMonth = 28
		}
	case 4, 6, 9, 11:
		daysInMonth = 30
	}

	day := rand.Intn(daysInMonth) + 1
	return year, month, day
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Versuch Nummer %v: \n", i+1)
		year, month, day := getRandomDate()
		fmt.Println(era, year, month, day)
		fmt.Printf("------------------\n")
		time.Sleep(400 * time.Millisecond)
	}
}

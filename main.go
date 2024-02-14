package main

import (
	"fmt"
	"math/rand"
)

func main() {

	// for loop needs to print every time all of the four attributes, so
	fmt.Printf("%-15v%-25v%-15v%-15v$ %-15v\n", 1, "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("2 =============================================================================")
	for i := 3; i < 13; i++ {
		fmt.Printf("%-15v%-25v%-15v%-15v$ %-15v\n", i, getSpaceCompany(), getTraveltime(), getTripType(), getPrice())
	}

}

func getSpaceCompany() string {
	choiceOne := "Virgin Galaxy"
	choiceTwo := "SpaceX"
	choiceThree := "Space Adventures"

	switch random := rand.Intn(3); random {
	case 0:
		return choiceOne
	case 1:
		return choiceTwo
	default:
		return choiceThree
	}
}

func getTraveltime() int {
	// get Traveltime to Mars at October 13th, 2020 in days
	return 62100000 / (rand.Intn(15) + 16) / 60 / 24
}

func getTripType() string {
	choiceOne := "Round-trip"
	choiceTwo := "One-way"

	switch random := rand.Intn(3); random {
	case 0:
		return choiceOne
	default:
		return choiceTwo
	}
}

func getPrice() int {
	return (rand.Intn(25) + 36) * 1000000
}

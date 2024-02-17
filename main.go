package main

import "fmt"

func main() {
	const lightYearInKilometers = 9.461e+12
	const distanceInKilometers = 236000000000000000
	const distanceInLightYears = distanceInKilometers / lightYearInKilometers
	fmt.Printf("The Canis Major Dwarf Galaxy is %.2f lightyeas aways from us.", distanceInLightYears)
}

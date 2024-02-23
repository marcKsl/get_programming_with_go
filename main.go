package main

import (
	"fmt"
)

func kelvinToCelcius(k float64) float64 {
	return k - 273.15
}

func celciusToFahrenheit(c float64) float64 {
	return (c * 9.0 / 5.0) + 32.0
}

func kelvinToFahrenheit(k float64) float64 {
	return celciusToFahrenheit(kelvinToCelcius(k))
}

func main() {
	k := 300.00

	fmt.Println(kelvinToCelcius(k))
	fmt.Println(kelvinToFahrenheit(k))

	c := 20.0

	fmt.Println(celciusToFahrenheit(c))
}

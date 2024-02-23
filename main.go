package main

import (
	"fmt"
)

type celsius float64
type kelvin float64

func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}

// 127° C, the surface temperature of the sunlit moon
func main() {
	var t celsius = 127
	var tk kelvin = celsiusToKelvin(t)
	fmt.Printf("If the sun lights the moon, its surface is %v° C, or respectively %v°K hot.", t, tk)
}

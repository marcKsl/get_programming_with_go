package main

import (
	"fmt"
)

// Two Tables, Celsius-Fahrenheit | Fahrenheit-Celsius
// Each shows the values for -40° C to 100° C in 5° C steps and respectively for fahrenheit
// Implement a function drawTable, that accepts a function as a parameter
// The accepted function sets the data for each row of the table.

// Save reusable code fragements in constants
const (
	spacer = "==========================="
	row    = "|   %-8s |   %-8s |\n"
)

type celsius float64
type fahrenheit float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c*5/9 + 32)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) * 9 / 5)
}

func displayCelsiusData(i int) (string, string) {
	d := celsius(i)
	c := fmt.Sprintf("%.1f", d)
	f := fmt.Sprintf("%.1f", d.fahrenheit())
	return c, f
}

func displayFahrenheitData(i int) (string, string) {
	d := fahrenheit(i)
	f := fmt.Sprintf("%.1f", d)
	c := fmt.Sprintf("%.1f", d.celsius())
	return f, c
}

func drawTable(d func(i int) (string, string)) {
	fmt.Println(spacer)
	fmt.Printf(row, "°C", "°F")
	fmt.Println(spacer)
	for i := -40; i < 101; i += 5 {
		a, b := d(i)
		fmt.Printf(row, a, b)
	}
	fmt.Println(spacer)
}

func main() {
	drawTable(displayCelsiusData)
	fmt.Printf("\n\n\n")
	drawTable(displayFahrenheitData)
}

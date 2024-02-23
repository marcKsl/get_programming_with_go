package main

import "fmt"

type celsius float64
type fahrenheit float64
type kelvin float64

func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit(c*5/9 + 32)
}

func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func (f fahrenheit) celsius() celsius {
	return celsius((f - 32) * 9 / 5)
}

func (f fahrenheit) kelvin() kelvin {
	return f.celsius().kelvin()
}

func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}

func (k kelvin) fahrenheit() fahrenheit {
	return k.celsius().fahrenheit()
}

func main() {

	var c celsius = 20
	var f fahrenheit = 43.11
	var k kelvin = 293.43

	fmt.Println(c)
	fmt.Println(c.fahrenheit())
	fmt.Println(c.kelvin())
	fmt.Println("----------------------------------------")
	fmt.Println(f)
	fmt.Println(f.celsius())
	fmt.Println(f.kelvin())
	fmt.Println("----------------------------------------")
	fmt.Println(k)
	fmt.Println(k.celsius())
	fmt.Println(k.fahrenheit())
	fmt.Println("----------------------------------------")
}

package main

import (
	"fmt"
	"math"
)

func main() {
	// The following int16 will overflow
	// if the if clause would not catch exceeding values,
	// causing it to wrap around and print:
	// -32768. The lowest possible number,
	// respectively the next possible number.
	// This is standard behaviour in go,
	// as it always tends to wrap around in such situations
	var bh float64 = 32768
	// The following constant values of math package are untyped.
	// As discussed in Lesson 8, constants in go are present at compile
	// time, thus they can also be untyped.
	// They infer the type out of the situation they are used in
	// Constants can also be typed.
	// That's why we can compare a float64 value to the constants directly:
	if bh < math.MinInt16 || bh > math.MaxInt16 {
		fmt.Println("Convertion not possible, number would overflow the type int16.")
		fmt.Println("Stay within the limits: Lowest:", math.MinInt16, ", highest:", math.MaxInt16)
	} else {
		var h = int16(bh)
		fmt.Println(h)
	}
	// Question: What code will determine if the variable v
	// is within the range of an 8-bit unsigned integer?
	v := 255
	if v <= math.MaxUint8 && v >= 0 {
		fmt.Println(v, "is in range of an 8-bit unsigned integer.")
	} else {
		fmt.Println(v, "is not in range of an 8-bit unsigned integer.")
	}
}

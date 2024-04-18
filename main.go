package main

import (
	"fmt"
	"math"
)

// Bad idea, should use struct embedding
// type report struct {
// 	sol         int
// 	temperature temperature
// 	location    location
// }

// Actually a bad way, and it should be done with struct embedding
// func (r report) average() celcius {
// 	// Just refer the method of the child here, to gain access to it from parent.
// 	// I like the idea.
// 	return r.temperature.average()
// }

type report struct {
	sol int
	location
	temperature
}

type temperature struct {
	high, low celcius
}

// We can not declare two methods with the same name on different embedded structs, and refer to them from the base struct
// That would lead to an ambiguous selector error, because the go compiler can not know where he needs to forward the call to.
// Can be resolved by creating the same method on the top-level, or referencing it there, to always make clear, which method should be called
func (t temperature) average() celcius {
	return celcius((float64(t.high) - math.Abs(float64(t.low))) / 2)
}

type location struct {
	lat, long float64
	// Can not declare high or low as values, because it would create an ambiguous selector error
	// That's because the Go compiler doesn't know which attribute, either location or temperature, he should forward to
}

type celcius float64

func main() {
	bradburry := location{-4.5895, 137.4417}
	t := temperature{high: -1.0, low: -7.0}
	report := report{sol: 15, temperature: t, location: bradburry}
	fmt.Printf("%+v\n", report)
	fmt.Printf("a balmy %v° C\n", report.temperature.high)
	fmt.Printf("average %v° C\n", report.average())
}

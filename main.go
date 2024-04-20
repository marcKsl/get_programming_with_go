package main

import "fmt"

type location struct {
	lat, long coordinate
}

// This satisfies the Stringer interface of the fmt package,
// Therefore friends like Println and Sprintf can use it directly.
func (l location) String() string {
	return fmt.Sprintf("%s\n%s", l.lat.String(), l.long.String())
}

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) String() string {
	// %v prints the variable and %c prints the character
	return fmt.Sprintf("%v°%v'%v''%c", c.d, c.m, c.s, c.h)
}

func main() {
	cologneLatitude := coordinate{50, 56, 6.6228, 'N'}
	cologneLongitude := coordinate{6, 57, 11.1636, 'E'}
	cologne := location{cologneLatitude, cologneLongitude}
	fmt.Println(cologne)
}

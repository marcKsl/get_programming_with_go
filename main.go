package main

import (
	"fmt"
	"math"
)

type coordinate struct {
	d, m, s float64
	h       rune
}

type location struct {
	lat  float64
	long float64
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func newLocation(lat, long coordinate) location {
	return location{lat.decimal(), long.decimal()}
}

type world struct {
	radius float64
}

var mars = world{radius: 3389.5}

// distance calculation using the Spherical Law of Cosines.
func (w world) distance(p1, p2 location) float64 {
	s1, c1 := math.Sincos(rad(p1.lat))
	s2, c2 := math.Sincos(rad(p2.lat))
	clong := math.Cos(rad(p1.long - p2.long))
	return w.radius * math.Acos(s1*s2+c1*c2*clong)
}

// rad converts degrees to radians.
func rad(deg float64) float64 {
	return deg * math.Pi / 180
}

func main() {

	fmt.Println("-------------------------------------------------------")
	fmt.Println("------------Attaching methods to structs---------------")
	fmt.Println("-------------------------------------------------------")

	// Bradbury Landing: 4°35'22.2" S, 137°26'30.1" E
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}

	fmt.Println(lat.decimal(), long.decimal())

	fmt.Println("-------------------------------------------------------")
	fmt.Println("------------Constructor functions----------------------")
	fmt.Println("-------------------------------------------------------")

	fmt.Println(newLocation(lat, long))

	fmt.Println("-------------------------------------------------------")
	fmt.Println("------------The class alternative----------------------")
	fmt.Println("-------------------------------------------------------")

	lat2 := coordinate{102, 35, 22.2, 'N'}
	long2 := coordinate{88, 26, 30.12, 'W'}

	lat4 := coordinate{17, 35, 22.2, 'S'}
	long4 := coordinate{8, 26, 30.12, 'E'}

	myLocation2 := newLocation(lat2, long2)
	myLocation4 := newLocation(lat4, long4)

	fmt.Println(mars.distance(myLocation2, myLocation4))

	fmt.Println("-------------------------------------------------------")
	fmt.Println("------------experiment: landing.go---------------------")
	fmt.Println("-------------------------------------------------------")

	spirit := newLocation(coordinate{14, 34, 6.2, 'S'}, coordinate{175, 228, 21.5, 'E'})
	opportunity := newLocation(coordinate{66, 34, 6.2, 'S'}, coordinate{59, 228, 21.5, 'E'})
	curiosity := newLocation(coordinate{1, 34, 6.2, 'N'}, coordinate{16, 228, 21.5, 'W'})
	insight := newLocation(coordinate{140, 34, 6.2, 'S'}, coordinate{120, 228, 21.5, 'W'})

	fmt.Printf("%+0.2v\n%+0.2v\n%+0.2v\n%+0.2v\n", spirit, opportunity, curiosity, insight)

	fmt.Printf("%v\n", mars.distance(spirit, opportunity))

	a := []location{spirit, opportunity, curiosity, insight}

	for i := range a {
		for j := range a {
			fmt.Println(mars.distance(a[i], a[j]))
		}
	}

	earth := world{radius: 6371.0}

	for i := range a {
		for j := range a {
			fmt.Println(earth.distance(a[i], a[j]))
		}
	}

}

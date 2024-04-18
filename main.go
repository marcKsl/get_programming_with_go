package main

import (
	"fmt"
	"math"
)

type rover struct {
	name string
	gps
}

type gps struct {
	currentLocation     location
	destinationLocation location
	world
}

func (g gps) distance() float64 {
	return g.world.distance(g.currentLocation, g.destinationLocation)
}

func (g gps) message() {
	fmt.Printf("%v kilometers remaining to the destination: %v", g.distance(), g.destinationLocation.name)
}

type location struct {
	name      string
	lat, long float64
}

func (l location) description() string {
	return fmt.Sprintf("Location: %v: %v, %v", l.name, l.lat, l.long)
}

type world struct {
	name   string
	radius float64
}

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
	bradburryLanding := location{lat: -4.5895, long: 137.4417}
	elysiumPlanitia := location{lat: 4.5, long: 135.9}
	mars := world{name: "Mars", radius: 3389.5}
	gps := gps{currentLocation: bradburryLanding, destinationLocation: elysiumPlanitia, world: mars}
	curiosity := rover{name: "Curiosity", gps: gps}
	curiosity.message()
}

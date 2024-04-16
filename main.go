package main

import (
	"fmt"
	"strings"
)

type Planets []string

func (p Planets) toStringWithSpace() string {
	return strings.Join(p, "--<-->--")
}

func (p Planets) terraform() Planets {
	var newPlanets Planets
	for _, planet := range p {
		newPlanets = append(newPlanets, "New "+planet)
	}
	return newPlanets
}

func main() {

	planets := Planets{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}

	fmt.Println(planets.terraform().toStringWithSpace())

}

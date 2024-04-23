package main

import "fmt"

type item string

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickup(i *item) {
	if i == nil {
		return
	}
	c.leftHand = i
	fmt.Printf("%v hat %v aufgehoben.\n", c.name, *i)
}

func (c *character) give(to *character) {
	if to == nil {
		return
	}
	if c.leftHand == nil {
		return
	}
	to.leftHand = c.leftHand
	c.leftHand = nil
	fmt.Printf("%v hat %v ein %v gegeben.", c.name, to.name, *to.leftHand)
}

func main() {
	sword := item("Schwert")
	arthur := character{name: "Arthur"}
	arthur.pickup(&sword)
	arthur.pickup(nil)
	fmt.Println(arthur)
	knight := character{name: "Knight"}
	fmt.Println(knight)
	arthur.give(&knight)
	fmt.Println(arthur)
	fmt.Println(knight)
}

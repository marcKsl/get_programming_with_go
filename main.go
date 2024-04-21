package main

import (
	"fmt"
)

type stats struct {
	level             int
	endurance, health int
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 + s.level)
	s.health = 5 * s.endurance
}

type character struct {
	name  string
	stats stats
}

type person struct {
	name, superpower string
	age              int
}

func (p *person) birthday() {
	p.age++
}

func birthday(p *person) {
	p.age++
}

func (p person) fakeBirthday() {
	p.age++
}

func main() {
	rebecca := person{
		name:       "Rebecca",
		superpower: "imagination",
		age:        14,
	}

	birthday(&rebecca)

	fmt.Printf("%+v\n", rebecca)

	// Same for methods
	nathan := &person{
		name: "Nathan",
		age:  17,
	}

	nathan.birthday()
	fmt.Printf("%+v\n", *nathan)

	// Even this works, because methods will automatically dereference the adress for the pointer receiver
	terry := person{
		name: "Terry",
		age:  15,
	}
	// We can just call birthday
	terry.birthday()
	fmt.Printf("%+v\n", terry)

	// But it is important that, the receiver of the birthday() method needs to be a pointer,
	// otherwise age really would not increment, because it would operate on a copy
	terry.fakeBirthday()
	fmt.Printf("%+v\n", terry)

	// interior pointers point to inner items of structs
	player := character{name: "Matthias"}
	levelUp(&player.stats)
	fmt.Printf("player is: %+v\n", player)
	playerStatAdress := &player.stats
	levelUp(playerStatAdress)
	playerStatAdressValue := *playerStatAdress
	fmt.Printf("player.stats after one more birthday is: %+v\n", player.stats)
	fmt.Printf("playerStatAdressValue is: %+v\n", playerStatAdressValue)
	levelUp(&player.stats)
	fmt.Printf("playerStatAdressValue after one original player levelup is: %+v\n", playerStatAdressValue)
	fmt.Printf("player after two original player levelup is: %+v\n", player.stats)

}

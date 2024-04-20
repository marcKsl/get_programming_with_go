package main

import (
	"fmt"
	"math/rand"
	"time"
)

// This means that everyone that is an Animaler shall have a String() Method that returns a string
// and an Eat() method that returns a string,
// aswell as a Move() method that also returns a string.°
type Animaler interface {
	String() string
	Eat() string
	Move() string
}

type lion struct {
	name string
}

func (l lion) String() string {
	return fmt.Sprintf("%s", l.name)
}

func (l lion) Eat() string {
	n := rand.Intn(3)
	switch n {
	case 0:
		return "I ate a good ol' zebra."
	case 1:
		return "I ate a good ol' tiger."
	default:
		return "I ate a good ol' gazelle."
	}
}

func (l lion) Move() string {
	return "I moved not a single meter today!"
}

type zebra struct {
	name string
}

func (z zebra) String() string {
	return fmt.Sprintf("%s", z.name)
}

func (z zebra) Eat() string {
	n := rand.Intn(3)
	switch n {
	case 0:
		return "I ate a good ol' watermelon today."
	case 1:
		return "I ate a good ol' strawberry today."
	default:
		return "I ate good ol' gras today."
	}
}

func (z zebra) Move() string {
	return "I moved seven kilometers today!"
}

func main() {
	day := 72
	dayNightPeriod := 12
	runtime := day / dayNightPeriod

	lion := lion{"Charles"}
	zebra := zebra{"Charlotte"}

	fmt.Println("-----------------------------------------")
	for i := runtime; i >= 0; i-- {
		if i%2 == 0 {
			fmt.Printf("Day %v: \n", runtime-i)
			n := rand.Intn(2)
			switch n {
			case 0:
				fmt.Println(lion)
				j := rand.Intn(2)
				switch j {
				case 0:
					fmt.Println(lion.Eat())
				default:
					fmt.Println(lion.Move())
				}
			default:
				fmt.Println(zebra)
				j := rand.Intn(2)
				switch j {
				case 0:
					fmt.Println(zebra.Eat())
				default:
					fmt.Println(zebra.Move())
				}
			}
		} else {
			fmt.Println("It is night, all animals are sleeping")
		}
		fmt.Println("-----------------------------------------")
		time.Sleep(time.Millisecond * 1500)
	}

}

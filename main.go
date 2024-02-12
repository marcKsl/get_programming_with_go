package main

import (
	"fmt"
	"math/rand"
)

func main() {

	if num := rand.Intn(3); num == 0 {
		fmt.Println("jajajaj")
	} else if num == 1 {
		fmt.Println("jojojojo")
	} else {
		fmt.Println("nenene")
	}

}

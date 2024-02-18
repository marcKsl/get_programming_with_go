package main

import (
	"fmt"
	"time"
)

func main() {

	// Use ROT13 to cipher a message that contains unicode characters, and thus runes(int32)

	message := "Hola Estación Espacial Internacional"
	fmt.Println("")
	fmt.Println("")
	for _, r := range message {
		c := r
		if c >= 'a' && c <= 'z' { // Lowercase
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		} else if c >= 'A' && c <= 'Z' { // Uppercase
			c = c + 13
			if c > 'Z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}

	fmt.Println("")
	fmt.Println("")
	fmt.Print("Translating")
	for i := 0; i < 10; i++ {
		fmt.Print(".")
		time.Sleep(150 * time.Millisecond)
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("=> Hola Estación Espacial Internacional")
	fmt.Println("")
	fmt.Println("")

}

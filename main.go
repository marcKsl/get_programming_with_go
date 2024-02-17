package main

import "fmt"

// In Go you need variables of type rune (int32) to get some of the characters inside a UTF-8 String.
// ASCII digits all fit into a byte (uint8) but the extended UTF-8 Character set contains values, that
// exceed 256.

func main() {
	s := "Hello, 世界"
	// If used with a string, the range keyword returns two values
	// The index in bytes (can be used as decimal value aswell) at the current position in the loop
	// And the rune value at that specific position (rember runes are an alias to int32)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("---Using runes to access single characters inside strings-------")
	fmt.Println("----------------------------------------------------------------")
	for index, runeValue := range s {
		fmt.Printf("%-3d=%2c\n", index, runeValue)
	}
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("---you can leave the index, and just use a _ blank identifier---")
	fmt.Println("----------------------------------------------------------------")
	for _, runeValue := range s {
		fmt.Printf("%c\n", runeValue)
	}
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("---This only works with the range keyword because it also------")
	fmt.Println("---returns the unicode code point, that is of type rune.--------")
	fmt.Println("----------------------------------------------------------------")
}

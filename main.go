package main

import "fmt"

func main() {
	// ROT13 cipher, just moves every character 13 steps higher.
	message := "uv vagreangvbany fcnpr fgngvba"

	for i := 0; i < len(message); i++ {
		c := message[i]
		// first line generally checks if character is between 'a' and 'z', so no space for example
		if c >= 'a' && c <= 'z' {
			c += 13
			// Handles the overflow in the given 26 character range
			if c >= 'z' {
				c -= 26
			}

		}
		// No need to fill a new string / array, just print out every character inline with no space between, genius.
		fmt.Printf("%c", c)
	}
}

package main

import (
	"fmt"
)

func main() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"

	// The strings.Repeat function may come in handy.
	// Give it a try, but also complete this exercise without
	// importing any packages other than fmt to print the deciphered message.
	// Try this exercise using range in a loop and again without it.
	// Remember that the range keyword splits a string into runes,
	// whereas an index like keyword[0] results in a byte.
	// Tip You can only perform operations on values of the same type,
	// but you can convert one type to the other (string, byte, rune).
	// To wrap around at the edges of the alphabet,
	// the Caesar cipher exercise made use of a comparison.
	// Solve this exercise without any if statements by using modulus (%).

	// for loop through len(cipherText)

	extendedKeyword := ""
	repeatCount := len(cipherText) / len(keyword)

	for i := 0; i < repeatCount+1; i++ {
		extendedKeyword += keyword
	}

	// There is a solution that GPT suggested that involves slices
	// We didn't use slices yet,
	// so i stick with just adding 1 to repeatCount for simplicity

	var output string
	for i := 0; i < len(cipherText); i++ {
		shift := (cipherText[i] - extendedKeyword[i] + 26) % 26
		output += string('A' + shift)
	}
	fmt.Println(output)
}

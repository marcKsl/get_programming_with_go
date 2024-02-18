package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	// Summary / Notes for Lesson 9

	question := "¿Cómo estás"
	question2 := "?Como estas"

	fmt.Println(len(question), "bytes")
	fmt.Println(len(question2), "bytes")
	// utf8 package gives you some interesting methods to perform on UTF related objects like strings, runes
	// This method counts the number of runes inside a given string, whereas len() counts the number of bytes
	// If using unicode signs, utf8 encoded characters can overflow normal byte (uint8) types, hence go uses runes(int32)
	// See also the last chapter for more clarification on the topic: Runes and Bytes
	fmt.Println(utf8.RuneCountInString(question), "runes")
	fmt.Println(utf8.RuneCountInString(question2), "runes")

	// This method takes a string and returns two values:
	// A rune: The first unicode code point it encounters
	// The size: The size of the rune in bytes, for normal ASCII digits it will be 1, but for some of the advanced characters
	// in the unicode set it will be 2, 3 or 4.
	c, size := utf8.DecodeRuneInString(question)
	c2, size2 := utf8.DecodeRuneInString(question2)
	fmt.Printf("First rune: %c %v bytes\n", c, size)
	fmt.Printf("First rune: %c %v bytes\n", c2, size2)

	// In UTF-8 unicode characters are of variable length and can be 1-4 bytes long.
	// A few examples:
	// The pile of poo emoji: 💩 (U+1F4A9)
	// The Mahjong tile red dragon: 🀄 (U+1F004)
	// The domino tile horizontal-00-05: 🁣 (U+1F063)

	poo := "💩"
	mahjong := "🀄"
	domino := "🁣"

	fmt.Println("--------------------------------")
	fmt.Println(len(poo))
	fmt.Println(len(mahjong))
	fmt.Println(len(domino))

	// If you use single quotes to initialize a character it will be stored as a unicode code point (rune/int32)
	// If you use double quotes it will always infer the string type

	test1 := '💩'
	test2 := "💩"

	test3 := 'A'
	test4 := "A"

	fmt.Println("--------------------------------")
	fmt.Printf("Poo emoji in single quotes gets printed as: %v and is of type %[1]T\n", test1)
	fmt.Printf("Poo emoji in double quotes gets printed as: %v and is of type %[1]T\n", test2)
	fmt.Printf("Capital A in single quotes gets printed as: %v and is of type %[1]T\n", test3)
	fmt.Printf("Capital A in double quotes gets printed as: %v and is of type %[1]T\n", test4)
	fmt.Println("--------------------------------")

	// How many runes and bytes are in the english alphabet?
	// To count runes we again use the RuneCountInString method from the utf8 package
	alphabet := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(len(alphabet))
	fmt.Println(utf8.RuneCountInString(alphabet))
	fmt.Println("--------------------------------")
	// How many bytes are in '¿'?
	openingQuestionMark := "¿"
	fmt.Println(len(openingQuestionMark))
	fmt.Println("--------------------------------")

	// You can use raw strings, already mentioned in previous lessons to ignore all escaping
	rawStringy := `\n this will be printed, all of it \\\\\\\n yeahhh\n\n, yeah.`
	fmt.Println(rawStringy)
	fmt.Println("--------------------------------")

	// Strings are immutable, single characters can be accessed but not altered
	myStringy := "Hello"
	for i := 0; i < len(myStringy); i++ {
		fmt.Printf("%c", myStringy[i])
		// this is not possible
		// myStringy[2] = "i"
	}
	fmt.Printf("\n")
	fmt.Println("--------------------------------")

	// The range keyword can decode a UTF-8 encoded string into runes
	// It gives us two values, the current index, and the rune (unicode code point) at that index
	myStringFullOfPotentialRunes := "Oha"
	for i, c := range myStringFullOfPotentialRunes {
		fmt.Printf("At position %v is the character %c\n", i, c)
	}
	fmt.Println("--------------------------------")
}

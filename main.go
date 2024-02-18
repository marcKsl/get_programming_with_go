package main

import (
	"fmt"
	"strconv"
)

func main() {
	var pi rune = 960
	var alpha rune = 940
	var omega rune = 969
	var bang byte = 33

	// string() interpretes the unicode code point and renders it as a string (using utf-8 encoding)
	fmt.Println("---------------------------------------------------------------------------")
	fmt.Print(string(pi), string(alpha), string(omega), string(bang))
	fmt.Println("")
	fmt.Println("---------------------------------------------------------------------------")
	// strconv.Itoa converts an Integer directly to a string value (the value that you can see)
	// Itoa is short for integer to ascii, but it does span the whole range of unicode, no worries
	fmt.Print(strconv.Itoa(int(pi)), strconv.Itoa(int(alpha)), strconv.Itoa(int(omega)), strconv.Itoa(int(bang)))
	fmt.Println("")
	fmt.Println("---------------------------------------------------------------------------")

	countdown := 10

	fmt.Println("---------------------------------------------------------------------------")
	str := "Launch in T minus " + strconv.Itoa((countdown)) + " seconds."
	str2 := "Launch in T minus " + string(countdown) + " seconds."
	fmt.Println(str)
	fmt.Println(str2)

	// The unicode standard encompasses all prior character encoding schemes
	// That includes ASCII, which form the first 128 characters
	// Note that some characters are not printable, because they derive back to a time
	// in 1960 when teletype machines where used, like punch-card media. Those characters
	// encrypt control characters like the newline character \n or the return carriage character
	// \r
	// Note that they do not get printed tho, because many modern systems use placeholders or
	// not print anything at all. In windows you might see strange glyphs cause it uses Code Page 437
	// by default to render characters in the lower range, unlike mac and linux that use standard UTF-8

	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("Here come the first 128 Unicode characters:")
	fmt.Println("---------------------------------------------------------------------------")
	for i := 0; i < 129; i++ {
		fmt.Println(i, ":", string(i))
	}

	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("You can also use fmt.Sprintf() to return strings")
	fmt.Println("directly rather than displaying them.")
	fmt.Println("---------------------------------------------------------------------------")

	myString := fmt.Sprintf("Launch in T minus %v seconds.", countdown)
	fmt.Println(myString)

	fmt.Println("---------------------------------------------------------------------------")
	fmt.Println("To go back from a string to an Integer we can use the Atoi function.")
	fmt.Println("That is short for ASCII to Integer. But it my return errors.")
	fmt.Println("But it does span the whole range of unicode, no worries")
	fmt.Println("---------------------------------------------------------------------------")

	myNumberInAString := "69oopsie"
	myNumberOutofAString, err := strconv.Atoi(myNumberInAString)
	if err != nil {
		fmt.Println("Oh no, something went wrong!")
	} else {
		fmt.Println(myNumberOutofAString)
	}

	// Go is statically typed, to let the compiler optimize code more,
	// so that it can create faster programs.
	// There is no way to change the type of a set variable in go.
	// Unlike in other languages like Python, JavaScript or Ruby
	// Go can not resolve the following lines of code:
	// var countdown = 10
	// countdown = 0.5
	// countdown = fmt.Sprintf("%v seconds", countdown)
	// Careful, the above lines will not compile in Go.
}

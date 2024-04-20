package main

import "fmt"

func main() {
	var administrator *string
	scolese := "CJ Scolese"
	administrator = &scolese
	fmt.Println(*administrator)

	bolden := "CF Bolden"
	administrator = &bolden
	fmt.Println(*administrator)

	*administrator = "C Johnson"
	fmt.Println(bolden)

	major := administrator
	*major = "Major Tom"
	fmt.Println(bolden)

	lightfoot := "RM Lightfoot"
	administrator = &lightfoot
	// major still points to bolden, because it is a copy
	fmt.Println(administrator == major)

	charles := *major
	*major = "Charles Bolden"
	// This prints Major Tom, correct
	fmt.Println(charles)
	// This prints Charles Bolden, correct
	fmt.Println(bolden)

	charles = "Charles Bolden"
	fmt.Println(charles == bolden)
	fmt.Println(&charles == &bolden)

	type person struct {
		name, superpower string
		age              int
	}

	timmy := &person{
		name: "Timmothy",
		age:  10,
	}

	timmy2 := person{
		name: "Timmothy2",
		age:  10,
	}

	timmy.superpower = "flying dragons"
	(*timmy).superpower = "riding horses"

	timmy2.superpower = "flying dragons"
	// This doesn't work, because it is not a pointer, instead it is a real instance of person
	// (*timmy2).superpower = "riding horses"

	fmt.Printf("%v, type: %[1]T", timmy)
	fmt.Printf("%v, type: %[1]T", timmy2)

	nonsense := &[]int{10}
	fmt.Println(nonsense)

	superpowers := &[3]string{"flying", "invisibility", "super strength"}

	// If it is a slice this needs to be done, because slices do not seem to
	// dereference automatically, maps also not
	// But both can be prefixed with & to create a pointer, only with composite literal tho
	fmt.Println((*superpowers)[0])
	// This only works for arrays not for slices
	fmt.Println(superpowers[0])
}

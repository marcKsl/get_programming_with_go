package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type location struct {
	lat  float64
	long float64
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {

	fmt.Println("-------------------------------------------------------")
	fmt.Println("------------Declaring a structure----------------------")
	fmt.Println("-------------------------------------------------------")

	var curiosity struct {
		lat  float64
		long float64
	}

	curiosity.lat = -4.5895
	curiosity.long = 137.4417

	fmt.Println(curiosity.lat, curiosity.long)
	fmt.Println(curiosity)

	fmt.Println("-------------------------------------------------------")
	fmt.Println("------------Reusing structures with types--------------")
	fmt.Println("-------------------------------------------------------")

	type location struct {
		lat  float64
		long float64
	}

	var spirit location
	spirit.lat = -14.5684
	spirit.long = 175.472636

	var opportunity location
	opportunity.lat = -1.9462
	opportunity.long = 354.4734

	fmt.Println(spirit, opportunity)
	fmt.Printf("Type of spririt: %T\n", spirit)
	fmt.Printf("Type of opportunity: %T\n", opportunity)

	fmt.Println("-------------------------------------------------------")
	fmt.Println("-----Initialize structures with composite literals-----")
	fmt.Println("-------------------------------------------------------")

	type location2 struct {
		lat, long float64
	}

	opportunity2 := location{lat: -1.9462, long: 354.4734}
	fmt.Println(opportunity2)

	insight := location{lat: 4.5, long: 135.9}
	fmt.Println(insight)

	// It is also possible to omit the keys and only pass the values,
	// use with caution tho, it can get pretty unreadable

	newthing := location{-44.5544, 754.7574}
	fmt.Println(newthing)

	// %v format verb with + prints also the field names

	fmt.Printf("%v\n", newthing)
	fmt.Printf("%+v\n", newthing)

	fmt.Println("-------------------------------------------------------")
	fmt.Println("------------Structures are copied----------------------")
	fmt.Println("-------------------------------------------------------")

	bradbury := location{-4.5895, 137.4417}
	curiosity2 := bradbury

	myRover := bradbury

	fmt.Printf("My Rover: %+v\n", myRover)

	curiosity2.long += 0.0106

	fmt.Println(bradbury, curiosity2)

	// Structs passed into functions also get copied

	fmt.Println("-------------------------------------------------------")
	fmt.Println("------------A slice of structures----------------------")
	fmt.Println("-------------------------------------------------------")

	type location3 struct {
		name string
		lat  float64
		long float64
	}

	locations3 := []location3{
		{name: "Bradbury Landing", lat: -4.5895, long: 137.4417},
		{name: "Columbia Memorial Station", lat: -14.5684, long: 175.472636},
		{name: "Challenger Memorial Station", lat: -1.9462, long: 354.4734},
	}

	for i := range locations3 {
		fmt.Printf("%+v\n", locations3[i])
	}

	fmt.Println("-------------------------------------------------------")
	fmt.Println("------------Encoding structures to JSON----------------")
	fmt.Println("-------------------------------------------------------")

	// Keys need to be uppercase to be exported in Go,
	// so that the JSON package can access them.
	type location4 struct {
		Lat  float64 `json:"labernicht"`
		Long float64 `json:"dochisso"`
	}

	curiosity4 := location4{-4.5895, 137.4417}

	bytes, err := json.Marshal(curiosity4)
	exitOnError(err)

	fmt.Println(string(bytes))

	fmt.Println("-------------------------------------------------------")
	fmt.Println("-----------Customizing JSON with struct tags-----------")
	fmt.Println("-------------------------------------------------------")

	type location5 struct {
		// struct tags are strings appended to an entry containing key value pairs,
		// that often correspond to a specific package

		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	curiosity5 := location5{-4.5895, 137.4417}

	bytes2, err2 := json.Marshal(curiosity5)
	exitOnError(err2)

	fmt.Println(string(bytes2))

	fmt.Println("-------------------------------------------------------")
	fmt.Println("------------experiment landing.go----------------------")
	fmt.Println("-------------------------------------------------------")

	type location10 struct {
		Name string  `json:"name"`
		Lat  float64 `json:"latitude"`
		Long float64 `json:"longitude"`
	}

	locations10 := []location10{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}

	bytes10, err10 := json.MarshalIndent(locations10, "", "")
	exitOnError(err10)

	fmt.Println(string(bytes10))
}

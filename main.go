package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {

	fmt.Println("-------------------------------------------------------")
	fmt.Println("-------------------Declaring a map---------------------")
	fmt.Println("-------------------------------------------------------")

	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
	}

	temp := temperature["Earth"]
	fmt.Printf("On average the Earth is %v C°.\n", temp)

	temperature["Earth"] = 16
	temperature["Venus"] = 464

	fmt.Println(temperature)

	// If you try to access a not existing key in a map, it returns the zero value for this type
	// Which, in this case, is 0, because we chose int as the type for our values

	fmt.Println(temperature["Moon"])

	// Go provides a comma, ok syntax for maps:
	// moon and ok are just example names, they can be anything

	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %v C°.\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}

	fmt.Println("-------------------------------------------------------")
	fmt.Println("-------------------Maps aren't copied------------------")
	fmt.Println("-------------------------------------------------------")

	planets := map[string]string{
		"Earth": "Sector ZZ9",
		"Mars":  "Sector ZZ9",
	}

	fmt.Println(planets["Earth"])

	planetsMarkII := planets
	planets["Earth"] = "whoops"

	fmt.Println(planets["Earth"])
	fmt.Println(planetsMarkII["Earth"])

	delete(planets, "Earth")
	fmt.Println(planetsMarkII)

	fmt.Println("-------------------------------------------------------")
	fmt.Println("-----------Preallocating maps with make----------------")
	fmt.Println("-------------------------------------------------------")

	temperatures := make(map[float64]int, 8)
	fmt.Println(temperatures)
	temperatures[4.0] = 77
	fmt.Println(temperatures)

	fmt.Println("-------------------------------------------------------")
	fmt.Println("------------Using maps to count things-----------------")
	fmt.Println("-------------------------------------------------------")

	temperaturez := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	frequency := make(map[float64]int)

	for _, t := range temperaturez {
		frequency[t]++
	}

	for t, num := range frequency {
		fmt.Printf("%+.3f occurs %d times\n", t, num)
	}

	fmt.Println("-------------------------------------------------------")
	fmt.Println("----------Grouping data with maps and slices-----------")
	fmt.Println("-------------------------------------------------------")

	temperatures2 := []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	groups := make(map[float64][]float64)

	for _, t := range temperatures2 {
		g := math.Trunc(t/10) * 10
		groups[g] = append(groups[g], t)
	}

	for g, temperatures := range groups {
		fmt.Printf("%+.0f: %v\n", g, temperatures)
	}

	fmt.Println("-------------------------------------------------------")
	fmt.Println("-------------Repurposing maps as sets------------------")
	fmt.Println("-------------------------------------------------------")

	var temperatures3 = []float64{
		-28.0, 32.0, -31.0, -29.0, -23.0, -29.0, -28.0, -33.0,
	}

	// This for loop ensures that every member in the set is unique, because thats what sets are,
	// unique collections of things

	set := make(map[float64]bool)
	for _, t := range temperatures3 {
		set[t] = true
	}

	if set[-28.0] {
		fmt.Println("set member")
	}

	if set[33.0] {
		fmt.Println("set member")
	} else {
		fmt.Println("Not there.")
	}

	fmt.Println(set)

	fmt.Println("-------------------------------------------------------")
	fmt.Println("----------------experiment words.go--------------------")
	fmt.Println("-------------------------------------------------------")

	// Question:

	// 	Write a function to count the frequency of words in a string of text and return a map of words with their counts. The function should convert the text to lowercase, and punctuation should be trimmed from words. The strings package contains several helpful functions for this task, including Fields, ToLower, and Trim.
	// Use your function to count the frequency of words in the following passage and then display the count for any word that occurs more than once.

	// Solution:

	// in: string
	// out: map[word string]count int

	// strings.Fields() can be used to get the string slice array

	var countFrequency = func(s string) map[string]int {
		inputBuffer := strings.Fields(s)
		wordSlices := make([]string, len(inputBuffer))
		for i := range inputBuffer {
			wordSlices[i] = strings.ToLower(strings.Trim(inputBuffer[i], " "))
		}
		buffer := make(map[string]int, 100)
		result := make(map[string]int, 100)
		for _, word := range wordSlices {
			buffer[word] = buffer[word] + 1
		}

		for k, v := range buffer {
			if v > 1 {
				result[k] = v
			}
		}

		return result
	}

	fmt.Println(countFrequency("As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."))

}

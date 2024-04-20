package main

import (
	"encoding/json"
	"fmt"
)

type location struct {
	lat, long coordinate
}

// This satisfies the Stringer interface of the fmt package,
// Therefore friends like Println and Sprintf can use it directly.
func (l location) String() string {
	return fmt.Sprintf("%s\n%s", l.lat.String(), l.long.String())
}

type coordinate struct {
	d, m, s float64
	h       rune
}

func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func (c coordinate) MarshalJSON() ([]byte, error) {
	// Calculate decimal degrees
	decimal := c.decimal()

	// We use a map, to create the existing fields.
	jsonData := map[string]interface{}{
		"decimal":    decimal,
		"dms":        c.String(),
		"degrees":    int(c.d),
		"minutes":    int(c.m),
		"seconds":    int(c.s),
		"hemisphere": string(c.h),
	}
	return json.Marshal(jsonData)
}

func (c coordinate) String() string {
	// %v prints the variable and %c prints the character
	return fmt.Sprintf("%v°%v'%v''%c", c.d, c.m, c.s, c.h)
}

func main() {
	coord := coordinate{
		d: 135,
		m: 54,
		s: 0,
		h: 'E',
	}

	// Marshal the coordinate instance to JSON
	jsonData, err := json.Marshal(coord)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the JSON data
	fmt.Println(string(jsonData))
}

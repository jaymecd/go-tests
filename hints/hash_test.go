package hash_test

import (
	"errors"
	"fmt"
)

var timeZone = map[string]int{
	"UTC": 0,
	"EST": -2,
	"CST": -6,
	"MST": -7,
	"PST": -8,
}

func fetchTZ(tz string) (int, error) {
	if seconds, ok := timeZone[tz]; ok {
		return seconds, nil
	}

	err := errors.New("timezone not found")
	var sec int

	return sec, err
}

func ExampleReverse() {
	table := []string {
		"UTC",
		"EST",
		"MST",
		"MSK",
	}

	for _, tz := range table {
		sec, err := fetchTZ(tz)

		if err != nil {
			fmt.Printf("Error: %s - %s\n", tz, err)
		} else {
			fmt.Printf("%s offset: %d\n", tz, sec)
		}
	}

	// Output: 
	// UTC offset: 0
	// EST offset: -2
	// MST offset: -7
	// Error: MSK - timezone not found
}

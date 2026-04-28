package main

import (
	"fmt"
)

func main() {
	//this function check how many times a word appears in a slice and record it in a map

	words := []string{"map", "go", "map", "go", "map", "staff"}

	count := make(map[string]int) // an empty map with a key of type string and a value of type int

	for _, word := range words {

		if word == "map" { // if the loop finds map it save it in the map and how many time it was found
			count["map"]++
		}
		if word == "go" {
			count["go"]++
		}
		if word == "staff" {

		}
	}
	fmt.Println(count)

}


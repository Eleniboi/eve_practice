package main

import "fmt"

func render(input string, lines []string) {

	for row := 1; row <= 8; row++ {
		for _, r := range input {

			if r < 32 || r > 126 {
				continue
			}
			index := int(r - 32)
			start := index * 9
			if start+row < len(lines) {

				fmt.Print(lines[start+row])
			}
		}
		fmt.Println()
	}
}

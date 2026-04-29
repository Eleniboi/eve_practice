package main

import (
	"fmt"
	"strings"
)

func render(input string, lines []string) {

	result := strings.Split(input, `\n`)
	
	for i := 0; i < len(result); i++ {
		for row := 0; row < 8; row++ {

			for _, ch := range result[i] {

				index := int(ch - 32)
				start := index * 9
				fmt.Println(lines[start+row])

			}
			fmt.Println()
		}
	}
}

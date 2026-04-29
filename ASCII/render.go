package main

import (
	"fmt"
	"os"
	"strings"
)

func render(lines []string) string {

	input := os.Args[1]

	if input == "" {

		return ""
	}

	var Output string

	result := strings.Split(input, `\n`)

	for i := 0; i < len(result); i++ {

		if result[i] == "" {
			return result[i]
		}

		for row := 0; row < 8; row++ {

			for _, ch := range result[i] {

				index := int(ch - 32)
				start := index * 9
				Output = lines[start+row]

			}
			fmt.Println()
		}
	}
	return Output

}

package main

import (
	"strings"
)

func render(input string, banner map[rune][]string) []string {

	splitInput := strings.Split(input, `\n`)

	var result []string
	var build strings.Builder
	for _, word := range splitInput {

		for i := 0; i < 8; i++ {

			for _, r := range word {

				build.WriteString(banner[r][i])
			}
			result = append(result, build.String())
			build.Reset()
		}
	}
	return result
}

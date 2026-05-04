package main

import (
	"strings"
)

func render(input string, banner map[rune][]string) []string {

	slitInput := strings.Split(input, "\\n")

	var result []string
	var build strings.Builder
	for _, char := range slitInput {

		for i := 0; i < 8; i++ {

			for _, r := range char {

				build.WriteString(banner[r][i])
			}
			result = append(result, build.String())
			build.Reset()
		}

	}

	return result
}

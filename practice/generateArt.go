package main

import (
	"strings"
)

func generateArt(input string, banner map[rune][]string) string {

	var build strings.Builder

	splitInput := split(input)

	for i := 0; i < len(splitInput); i++ {

		if splitInput[i] == "" {

			build.WriteString("\n")
			continue
		}
		result := render(splitInput[i], banner)

		for j := 0; j < len(result); j++ {

			build.WriteString(result[j])
			build.WriteString("\n")
		}
	}
	return build.String()
}

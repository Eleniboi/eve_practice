package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	input := os.Args[1]

	banner := os.Args[2]

	file, err := os.ReadFile(banner)

	if err != nil {
		fmt.Println("Error reading file")
		return
	}

	splitInput := strings.Split(input, "\\n")

	fileLines := strings.Split(strings.ReplaceAll(string(file), "\r\n", "\n"), "\n")
	//fileLines = fileLines[1:]

	var bannerFont = make(map[rune][]string)

	for i := ' '; i <= '~'; i++ {

		startindx := int(i-32)*9 + 1
		endindx := startindx + 8
		bannerFont[i] = fileLines[startindx:endindx]
	}

	for _, word := range splitInput {

		for i := 0; i < 8; i++ {

			for _, ch := range word {

				fmt.Print(bannerFont[ch][i])
			}
			fmt.Println()
		}
	}
}

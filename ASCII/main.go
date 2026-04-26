package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	input := os.Args[1]
	bannerFont := os.Args[2]

	file, err := os.Open(bannerFont)

	if err != nil {
		fmt.Println("Error opening file, ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}

	result := strings.Split(input, `\n`)

	for i := 0; i < len(result); i++ {

		if result[i] == "" {
			fmt.Println()
			continue
		}

		for row := 1; row <= 8; row++ {

			for _, r := range result[i] {
				index := int(r - 32)

				start := index * 9

				fmt.Print(lines[start+row])
			}
			fmt.Println()
		}
	}
}

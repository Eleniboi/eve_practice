package main

import (
	"bufio"
	"fmt"
	"os"
)

func LoadBanner(bannerFont string) []string {

	file, err := os.Open(bannerFont)

	if err != nil {
		fmt.Println("Error opening file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Error Scanner file")
	}

	return lines
}

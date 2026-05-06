package main

import (
	"fmt"
	"os"
	"strings"
)

func loadBanner(bannerFont string) (map[rune][]string, error) {

	data, err := os.ReadFile(bannerFont)

	if err != nil {
		return nil, fmt.Errorf("Error reading file %w", err)
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("oops! this file is empty")
	}
	fileLines := strings.Split(strings.ReplaceAll(string(data), "\r\n", "\n"), "\n")

	fileLines = fileLines[1:]
	var banner = make(map[rune][]string)
	for i := ' '; i <= '~'; i++ {

		startindx := int(i-32) * 9
		endindx := startindx + 8
		banner[i] = fileLines[startindx:endindx]
	}
	return banner, nil
}

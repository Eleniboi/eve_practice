package main

import (
	"fmt"
	"os"
	"strings"
)

func LoadBanner(bannerfile string) (map[rune][]string, error) {

	file, err := os.ReadFile(bannerfile)

	if err != nil {
		return nil, fmt.Errorf("Error reading file: %w", err)
	}

	fileLines := strings.Split(string(file), "\n")

	fileLines = fileLines[1:]

	bannerFont := make(map[rune][]string)

	for i := ' '; i <= '~'; i++ {

		startindx := int(i-32) * 9
		endindx := startindx + 8

		bannerFont[i] = fileLines[startindx:endindx]
	}
	return bannerFont, nil
}

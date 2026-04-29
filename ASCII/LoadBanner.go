package main

import (
	"bufio"
	"log"
	"os"
)

func LoadBanner(bannerFont string) []string {

	bannerFont = os.Args[2]

	if len(os.Args) == 2 {
		bannerFont = "standard.txt"
	}
	if len(os.Args) == 0 {
		return os.Args
	}

	file, err := os.Open(bannerFont)

	if err != nil {
		log.Fatal()
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}
	return lines
}

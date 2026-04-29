package main

import (
	"bufio"
	"log"
	"os"
)

func LoadBanner(bannerFont string) []string {

	file, err := os.Open(bannerFont)

	if err != nil {
		log.Fatal()
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal()
	}

	return lines
}

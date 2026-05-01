package main

import (
	"errors"
	"os"
	"strings"
)

func Load_Banner(bannerFont string) (map[rune][]string, error) {

	file, err := os.ReadFile(bannerFont)

	if err != nil {
		return nil, errors.New("Error reading file")

	}

	text := strings.Split(string(file), "\n")

	text = text[1:]

	var banner = make(map[rune][]string)

	for ch := ' '; ch <= '~'; ch++ {
		start := ch - 32
		end := start + 8
		banner[ch] = text[start:end]
	}
	return banner, nil
}

// func LoadBanner(bannerFont string) []string {

// 	file, err := os.Open(bannerFont)

// 	if err != nil {
// 		fmt.Println("Error opening file")
// 	}

// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	var lines []string

// 	for scanner.Scan() {

// 		lines = append(lines, scanner.Text())
// 	}

// 	err = scanner.Err()
// 	if err != nil {
// 		fmt.Println("Error Scanner file")
// 	}

// 	return lines
// }

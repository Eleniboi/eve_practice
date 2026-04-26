package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var bannerFont string

	inputfile := os.Args[1]
	bannerFont = os.Args[2]

	file, err := os.Open(bannerFont)

	if err != nil {
		fmt.Println("Error opening file, ", err)
		return
	}

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}

	for row := 1; row < 9; row++ {

		for _, r := range inputfile {

			index := int(r - 32)
			start := index * 9

			fmt.Print(lines[start+row])
		}
		fmt.Println()
	}

	//fmt.Println("samuel\n\nhow")
}

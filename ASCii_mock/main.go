package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	input := os.Args[1]

	banner := "standard.txt"

	result, _ := LoadBanner(banner)

	art := render(input, result)

	for i := 0; i < 8; i++ {

		time.Sleep(time.Second)

		fmt.Print(art)
		fmt.Println()

	}
}

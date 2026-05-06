package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	input := os.Args[1]

	bannerFont := os.Args[2]

	banner, err := LoadBanner(bannerFont)

	if err != nil {
		fmt.Println("Error loading banner")
		return
	}

	result := render(input, banner)

	for i := 0; i < len(result); i++ {

		time.Sleep(time.Second)

		fmt.Println(result[i])
	}
}

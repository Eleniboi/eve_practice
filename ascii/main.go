package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	input := os.Args[1]

	banner := "standard.txt"

	bannerFont, _ := LoadBanner(banner)

	result := render(input, bannerFont)

	//for _, ch := range input {

	for i := 0; i < 8; i++ {
		time.Sleep(time.Second)

		fmt.Print(result[i])
		fmt.Println()
	}
	//}

}

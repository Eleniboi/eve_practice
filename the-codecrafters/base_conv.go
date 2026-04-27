package main

import (
	"bufio"
	"fmt"
	"os"
)

func help() {
	fmt.Println("For Smooth Experience enter the following")
	fmt.Println("convert  1E hex  ---> 30")
	fmt.Println("convert  1010 bin ---> 10")
	fmt.Println("Enter help for this same menu")
	fmt.Println("Enter quit to exit the program")

}

func base() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("welcome to base converter")
	fmt.Print(`Enter help for Guildline or quit to exit the program: `)
	//fmt.Print(": ")
	for scanner.Scan() {

		fmt.Print(" > ")
		input := scanner.Text()

		if input == "" {
			fmt.Println("Enter help for Guildline")
			continue
		}
		if input == "help" {
			help()
			continue
		}
		if input == "quit" {
			fmt.Println("GoodBye!!")
			break
		}
	}
}

func main() {
	base()
}

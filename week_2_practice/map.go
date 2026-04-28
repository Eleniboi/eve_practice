package main


import (
	"fmt"
)

func main() {

	inventory := map[string]int{

		"Apples":10,
		"Banana":5,
		"cherries":0,
	}

	shopinglist := "cherries"

	quantity, exist := inventory[shopinglist]

	if !exist {
		fmt.Printf("we don't sell %v here", shopinglist)
	}else if quantity == 0{
		fmt.Printf("%v out of stock", exist)
	}else {
		fmt.Printf("still have %d %v left", quantity, exist)
	}
}
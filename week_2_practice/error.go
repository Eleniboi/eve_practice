package main

import (
	"errors"
	"fmt"
)

// this function recieves two numbers of type float64, float64 and error it also handle error like dividing a number by 0
func divide(a, b float64) (float64, error) {

	var result float64

	if b == 0 {
		return 0, errors.New("can not be divided by 0")
	}
	
	result = a / b
	//frff
	return result, nil
}

func main() {
	res, err := divide(10, 0)

	if err != nil {
		fmt.Println("Invalid, ",err)
		return
	}

	fmt.Println("result",res)
}

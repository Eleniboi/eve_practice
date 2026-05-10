package main

import "fmt"

func validate(input string) (rune, error) {

	var res rune
	for _, ch := range input{

		if ch >= ' ' && ch <= '~'{
			return ch, nil
		}

	}
	return 0, fmt.Errorf("%c is not a printable character", res)
}


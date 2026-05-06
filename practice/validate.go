package main

import "fmt"

func validate(input string) (rune, error) {

	for _, ch := range input {

		if ch >= ' ' && ch <= '~' {
			return ch, nil
		} else {
			return 0, fmt.Errorf(", %c is not a printable character", ch)
		}
	}
	return 0, nil
}

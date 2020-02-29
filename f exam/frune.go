package main

import (
	"github.com/01-edu/z01"
)

func FirstRune(s string) rune {

	var c rune

	for i, v := range s {
		if i == 0 {
			c = v
		}
	}
	return c

}

func main() {
	s := "jekg"
	z01.PrintRune(FirstRune((s)))
}

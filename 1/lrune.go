package main

import "github.com/01-edu/z01"

func LastRune(s string) rune {
	var c rune
	for i, r := range s {
		if i == len(s)-1 {
			c = r
		}
	}
	return c
}

func main() {
	s := "jekg"
	z01.PrintRune(LastRune((s)))
}

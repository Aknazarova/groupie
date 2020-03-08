package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 2 {
		z01.PrintRune(10)
	} else {
		str1 := []rune(arg[0])
		str2 := []rune(arg[1])
		count := 0
		for _, c := range str2 {

			if count == len(str1) {
				z01.PrintRune('1')
				z01.PrintRune(10)
				return

			}
			if str1[count] == c {
				count++
			}

		}
		z01.PrintRune('0')
		z01.PrintRune(10)

	}
}

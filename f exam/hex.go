package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		z01.PrintRune(10)
	} else {
		str := ""
		hex := "0123456789abcdef"
		rem := 0
		num, _ := strconv.Atoi(arg[0])

		for num > 0 {
			rem = num % 16
			str = string(hex[rem]) + str
			num /= 16
		}

		for _, c := range str {
			z01.PrintRune(c)
		}
	}
}

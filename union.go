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
		a := arg[0] + arg[1]

		ftext := ""
		flag := false

		for _, c := range a {
			flag = true
			for _, v := range ftext {
				if c == v {
					flag = false
				}
			}
			if flag == true {
				ftext += string(c)
			}
		}

		for _, b := range ftext {
			z01.PrintRune(b)
		}

	}
	z01.PrintRune(10)
}

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 2 {
		z01.PrintRune(10)
	} else {
		a, _ := strconv.Atoi(arg[0])
		b, _ := strconv.Atoi(arg[1])

		if a > b {
			a, b = b, a
		}
		fmt.Println(a)
		gcd := 0
		for i := a; i > 0; i-- {
			if a%i == 0 && b%i == 0 {
				gcd = i
				break
			}

		}
		fmt.Println(gcd)

	}
}

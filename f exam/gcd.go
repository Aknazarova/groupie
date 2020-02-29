package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 2 {
		fmt.Println()
	} else {
		a, _ := strconv.Atoi(arg[0])
		b, _ := strconv.Atoi(arg[1])

		if a > b {
			a, b = b, a
		}
		gcd := 0
		for i := 1; i <= a; i++ {
			if a%i == 0 && b%i == 0 {
				gcd = i

			}
		}
		fmt.Println(gcd)
	}
}

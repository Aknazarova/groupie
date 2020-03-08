package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	if len(arg) != 1 {
		z01.PrintRune(10)
	} else {
		num, _ := strconv.Atoi(arg[0])
		for i := 1; i <= num; i++ {
			if Power(2, i) == num {
				fmt.Println("true")
				return
			}

		}
	}
	fmt.Println("false")
}

func Power(n, pow int) int {
	if pow == 0 {
		return 1
	} else if pow == 1 {
		return n
	}
	return n * Power(n, pow-1)
}

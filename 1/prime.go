package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1:]
	a := arg[0]
	if len(arg) != 1 {
		fmt.Println()
	} else {
		num, _ := strconv.Atoi(a)
		div := 2
		for div <= num {
			if num%div == 0 {
				print(div)
				if num == div {
					fmt.Println()
					return
				}
				fmt.Print("*")
				num /= div
				div = 1
			}
			div++
		}

	}
}

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1:]
	x, _ := strconv.Atoi(arg[0])

	div := 2
	if len(arg) != 1 {
		fmt.Println()
	} else {
		for div <= x {
			if x%div == 0 {
				print(div)
				if x == div {
					fmt.Println("")
					return
				}
				fmt.Print("*")
				x /= div
				div = 1

			}
			div++
		}

	}
}

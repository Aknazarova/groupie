package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1:]
	res := ""
	if len(arg) != 1 {
		fmt.Println()
	} else {

		num, _ := strconv.Atoi(arg[0])

		for num != 1 {
			if num%2 == 1 {
				res = "false"
				break
			}
			num /= 2
			res = "true"

		}
		fmt.Print(res)
		// z01.PrintRune(10)
	}

}

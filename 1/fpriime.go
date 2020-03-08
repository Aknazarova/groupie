package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		fmt.Println()
	} else {
		num, err := strconv.Atoi(arg[0])
		if err != nil {
			fmt.Println()
			return
		}
		if num == 1 { //Если приходит 1, то отвечаем 1
			fmt.Println()
			return
		}
		if num < 1 { //Если меньше 1, то выходим
			fmt.Println()
			return
		}

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
		// fmt.Println()
	}
}

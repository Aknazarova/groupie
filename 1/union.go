package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	str := arg[0] + arg[1]
	res := ""
	flag := false
	if len(arg) != 2 {
		fmt.Println()
	} else {
		for _, c := range str {
			flag = true
			for _, d := range res {
				if c == d {
					flag = false
				}
			}

			if flag == true {
				res += string(c)
			}
		}

		fmt.Println(res)
	}
}

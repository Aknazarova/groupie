package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	res := ""
	flag := false
	if len(arg) != 2 {
		fmt.Println()
	} else {
		for _, c := range arg[0] {
			for _, d := range arg[1] {
				if c == d {
					flag = true

					for _, v := range res {
						if c == v {
							flag = false
						}
					}

					if flag == true {
						res += string(c)
					}
				}
			}

		}
		fmt.Println(res)
	}
}

package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	a := arg[0]
	if len(arg) != 1 {
		fmt.Println()
	} else {
		res := ""
		for i := len(a) - 1; i > 0; i-- {
			if a[i] == ' ' {
				res += string(a[i+1:]) + " "
				a = a[0:i]
			}
		}
		res += a
		fmt.Println(res)
	}
}

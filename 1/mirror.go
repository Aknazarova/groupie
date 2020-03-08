package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	var res string
	if len(arg) != 1 {
		fmt.Println()
	} else {
		for _, c := range arg[0] {
			if c >= 97 && c <= 122 {
				res += string(122 - c + 97)
			} else if c >= 65 && c <= 90 {
				res += string(90 - c + 65)
			} else if c >= 32 && c <= 64 {
				res += string(c)
			}
		}
		fmt.Print(res)
	}

}

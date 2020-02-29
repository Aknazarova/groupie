package main

import (
	"fmt"
	"os"
)

func main() {
	var c1 [1000]bool
	var c2 [1000]bool
	var sum [1000]int
	arg := os.Args[1:]
	if len(arg) != 1 {
		fmt.Println()
	} else {
		a := arg[0]
		for _, i := range a {
			sum[i]++
		}

		for _, i := range a {
			if c1[sum[i]] == false && c2[i] == false {
				c1[sum[i]] = true
				c2[i] = true

			} else if (c1[sum[i]] == true && c2[i] == false) || (c1[sum[i]] == false && c2[i] == true) {
				fmt.Print(false)
				return
			}
		}
		fmt.Print(true)
	}

}

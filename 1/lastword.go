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
		// z01.PrintRune(10)
	} else {
		last1 := 0
		last2 := 0
		res := ""
		for i := len(a) - 1; i >= 0; i-- {
			if last1 == 0 && a[i] != ' ' {
				last1 = i + 1
			} else if last1 != 0 && a[i] == ' ' {
				last2 = i + 1
				break
			}

		}
		res = a[last2:last1]
		fmt.Println(last1)
		fmt.Println(last2)
		// z01.PrintRune(10)
		fmt.Println(res)
	}

}

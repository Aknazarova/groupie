package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		z01.PrintRune(10)

	} else {
		last1 := 0
		last2 := 0
		a := arg[0]
		res := ""
		for i := len(a) - 1; i >= 0; i-- {
			if last1 == 0 && a[i] != ' ' {
				last1 = i + 1
			} else if last1 != 0 && a[i] == ' ' {
				last2 = i + 1
				break
			}
		}
		str := a[last2:last1]
		fmt.Println(str)
		res = ubr(a[last2:])

		for _, c := range res {
			z01.PrintRune(c)
		}
		z01.PrintRune(10)
		// fmt.Println(res)

		// fmt.Println(last1)
		// fmt.Println(last2)
	}
}

func ubr(str string) string {
	for i := 0; i < len(str); i++ {
		if str[0] == ' ' {
			return ubr(str[1:])
		}
	}

	for i := len(str) - 1; i >= 0; i-- {
		if str[len(str)-1] == ' ' {
			return ubr(str[:len(str)-1])
		}

	}
	return str
}

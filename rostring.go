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
		str := ""
		res := ""
		a := ubr(arg[0])
		for i := 0; i <= len(a)-1; i++ {
			if a[i] == ' ' {
				str = a[:i]
				// fmt.Println(str)
				a = a[i+1:]
				break

			}

		}
		r := ""
		res = a + " " + str
		for i := 0; i <= len(res)-1; i++ {
			if res[i] == ' ' {
				r = ubr(res)
			}
		}
		fmt.Println(r)

	}
}

func ubr(str string) string {
	res := ""
	if str[0] == ' ' {
		return ubr(str[1:])
	}

	for i := len(str) - 1; i > 0; i-- {
		if str[len(str)-1] == ' ' {
			return ubr(str[:len(str)-1])
		}
	}

	for i := 0; i <= len(str)-1; i++ {
		if str[i] == ' ' && str[i+1] == ' ' {
			continue
		}
		res += string(str[i])
	}

	return res
}

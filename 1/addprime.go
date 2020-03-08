package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	a := arg[0]
	str := ""
	num := 0
	res := 0
	if len(arg) != 1 {
		z01.PrintRune(10)
	} else {

		for _, i := range a {
			num = int(i-48) + num*10 //atoi
		}

		// fmt.Print(num)

		for i := 2; i <= num; i++ {
			count := 0
			for j := 1; j <= i; j++ {
				if i%j == 0 {
					count++
				}
			}
			if count == 2 {
				res += i
			}
		}

		// fmt.Print(res)

		for res > 0 {
			str = string(res%10+48) + str //itoa
			res /= 10
		}

		for _, c := range str {
			z01.PrintRune(c)
		}
		// fmt.Print(num)

	}
}

package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		z01.PrintRune(10)
	} else {
		res := ""
		num, _ := strconv.Atoi(arg[0])
		// if err!= nil{
		// 	z01.PrintRune('0')
		// }

		for i := 1; i <= 9; i++ {
			res = Itoa(num * i)
			z01.PrintRune(rune(i%10) + 48)

			z01.PrintRune(' ')
			z01.PrintRune('*')
			z01.PrintRune(' ')
			for _, c := range arg[0] {
				z01.PrintRune(c)
			}
			z01.PrintRune(' ')
			z01.PrintRune('=')
			z01.PrintRune(' ')
			for _, v := range res {
				z01.PrintRune(v)
			}
			z01.PrintRune(10)
		}

	}
}

func Itoa(n int) string {
	str := ""
	// res := ""

	flag := false

	if n < 0 {
		n = -n
		flag = true
	}
	for n != 0 {
		str = string(n%10+48) + str
		n /= 10
	}
	if flag == true {
		str = "-" + str
	}
	return str
}

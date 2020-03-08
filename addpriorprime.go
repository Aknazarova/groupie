package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		z01.PrintRune(10)
	} else {
		n := arg[0]
		num := 0
		res := 0
		str := ""
		for _, c := range n {
			num = int(c-48) + num*10
			// fmt.Println(num)
		}
		res = Priorprime(num)
		// fmt.Println(res)

		for res > 0 {

			str = string(res%10+48) + str
			res /= 10
			// break
		}

		for _, c := range str {
			z01.PrintRune(c)
		}
	}
	z01.PrintRune(10)
}

func Priorprime(x int) int {
	res := 0

	for i := 2; i <= x; i++ {
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
	return res

}

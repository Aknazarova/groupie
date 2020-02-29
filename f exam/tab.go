package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]

	a := arg[0]
	if len(arg) != 1 {

		z01.PrintRune(10)
	} else {
		num, _ := strconv.Atoi(a)
		for i := 1; i <= 9; i++ {
			res := Itoa(num * i)
			z01.PrintRune(rune(i%10 + 48))
			z01.PrintRune(' ')
			z01.PrintRune('*')
			z01.PrintRune(' ')
			for _, c := range a {
				z01.PrintRune(c)
			}
			z01.PrintRune(' ')
			z01.PrintRune('=')
			z01.PrintRune(' ')
			for _, c := range res {
				z01.PrintRune(c)
			}
			z01.PrintRune(10)
		}

	}

}
func Itoa(num int) string {
	str := ""

	for num != 0 {
		str = string(num%10+48) + str
		num /= 10
	}
	return str
}

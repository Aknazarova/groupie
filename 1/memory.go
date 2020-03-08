package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

func main() {
	arr := [10]int{104, 101, 108, 108, 111, 16, 21, 42}
	PrintMemory(arr)
}

func PrintMemory(arr [10]int) {
	for i, c := range arr {
		hex := NbrBase(c, "0123456789abcdef")
		fmt.Print(hex)
		n := len(hex)
		for n <= 8 {

			if n == 4 {

				fmt.Print(" ")
			} else {
				fmt.Print("0")
			}

			n++
		}
		fmt.Print(" ")
		fmt.Print(" ")
		if (i+1)%4 == 0 {
			fmt.Println()
		}

	}
	fmt.Println()
	for _, c := range arr {
		if c >= 31 {
			z01.PrintRune(rune(c))
		} else {
			z01.PrintRune('.')
		}
	}
	fmt.Println()
}

func NbrBase(n int, base string) string {
	res := ""
	flag := false
	if n < 0 {
		flag = true
		n = -n
	}
	for n != 0 {
		res = string(base[n%len(base)]) + res
		n /= len(base)
	}

	if flag {
		res = "-" + res
	}
	return res
}

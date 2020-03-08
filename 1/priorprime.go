package main

import "fmt"

func Priorprime(x int) int {
	res := 0

	for i := 2; i < x; i++ {
		count := 0
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}

		if count == 2 {
			res += i
		}
		// fmt.Print(res)
	}
	return res
}

func main() {
	fmt.Println(Priorprime(7))
}

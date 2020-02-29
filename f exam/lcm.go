package main

import "fmt"

func main() {
	fmt.Println(Lcm(2, 7))
	fmt.Println(Lcm(0, 4))
}

func Lcm(first, second int) int {
	if first > second {
		first, second = second, first
	}
	// fmt.Print(first)
	lcm := 0

	for i := first * second; i >= second; i-- {
		if i%first == 0 && i%second == 0 {
			lcm = i

		}
		i /= second
	}
	return lcm
}

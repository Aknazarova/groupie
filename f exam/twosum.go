package main

import "fmt"

func TwoSum(nums []int, target int) []int {
	a := []int{0, 0}
	for i, c := range nums {
		for j, d := range nums {
			if c+d == target && i != j {
				a[0] = c
				a[1] = d
				return a
			}

		}
	}
	return nil
}

func main() {
	case1 := []int{1, 2, 3, 4}
	out := TwoSum(case1, 5)
	fmt.Println(out)
}

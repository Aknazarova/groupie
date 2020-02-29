package main

import "fmt"

func AtoiBase(s string, base string) int {
	res := 0

	for _, v := range base {
		if v == '-' || v == '+' {
			return 0
		}
	}
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if base[i] == base[j] {
				return 0
			}
		}
	}

	for _, c := range s {
		for i, v := range base {
			if c == v {
				res = res*len(base) + i
			}
		}
	}
	return res
}

func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}

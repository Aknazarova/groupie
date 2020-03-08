package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "HelloHAhowHAareHAyou?"
	fmt.Println(Split(str, "HA"))
	fmt.Println(strings.Split(str, "HA"))
}

func Split(str, charset string) []string {
	last := 0
	var res []string
	for i := 0; i < len(str); i++ {
		if i+len(charset) <= len(str) {
			if str[i:i+len(charset)] == charset {
				fmt.Println("1: ", last, "2: ", i)
				res = aappend(res, str[last:i])
				last = i + len(charset)
				i = i + len(charset) - 1
			}
		}

	}
	res = aappend(res, str[last:])
	return res
}

func aappend(arr []string, s string) []string {
	newArr := make([]string, len(arr)+1)
	for i, w := range arr {
		newArr[i] = w
	}

	newArr[len(arr)] = s
	return newArr
}

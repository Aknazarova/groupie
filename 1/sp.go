package main

import (
	"fmt"
)

func main() {
	str := "HelloHAhowHAareHAyou?"
	fmt.Println(Split(str, "HA"))
}

func Split(str, charset string) []string {
	var res []string
	last := 0
	for i := 0; i < len(str); i++ {
		if i+len(charset) <= len(str) {
			if str[i:i+len(charset)] == charset {
				res = aappend(res, str[last:i])
				last = i + len(charset)
				i = i + len(charset) - 1
			}
		}
	}
	res = aappend(res, str[last:])
	return res

}

func aappend(str []string, s string) []string {
	new := make([]string, len(str)+1)

	for i, c := range str {
		new[i] = c
	}
	new[len(str)] = s
	return new
}

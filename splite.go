package main

import "fmt"

func main() {
	str := "HAHelloHAhowHAareHAyou?"
	fmt.Println(Split(str, "HA"))
}

func Split(str, ch string) []string {
	index := 0
	var res []string
	for i := 0; i < len(str)-1; i++ {
		if len(ch) <= len(str) {
			if str[i:i+len(ch)] == ch {
				res = Append(res, str[index:i])
				index = i + len(ch)
				i += len(ch) - 1
			}
		}

	}
	res = Append(res, str[index:])
	return res
}

func Append(str []string, s string) []string {
	new := make([]string, len(str)+1)

	for i, c := range str {
		new[i] = c
	}
	new[len(str)] = s
	return new
}

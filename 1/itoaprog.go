package main

import "fmt"

func Itoa(n int) string {
	str := ""
	flag := false
	if n < 0 {
		flag = true
		n = -n
	}

	for n != 0 {
		str = string(n%10+48) + str
		n /= 10
	}

	if flag == true {
		str = "-" + str
	}
	return str

}

func main() {
	fmt.Println(Itoa(125))
	fmt.Println(Itoa(-125))
}

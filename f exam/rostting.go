package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]

	if len(arg) != 1 {
		fmt.Println()
	} else {
		str := ""
		res := space(arg[0])
		fmt.Println(res)
		for i := 0; i < len(res); i++ {
			if res[i] == ' ' {
				str += res[:i]
				res = res[i+1:]
				break
			}
		}
		str = res + " " + str
		fmt.Println(str)
	}
}

func space(str string) string {
	res := ""
	if str[0] == ' ' {
		return space(str[1:])
	}

	n := len(str)
	if str[n-1] == ' ' {
		return space(str[:n-1])
	}
	for i := 0; i < n; i++ {
		if str[i] == ' ' && str[i+1] == ' ' {
			continue
		}
		res += string(str[i])
	}
	return res

}

// func main() {
// 	arg := os.Args[1:]

// 	if len(arg) != 1 {
// 		fmt.Println()
// 	} else {
// 		res := ubr(arg[0])
// 		// fmt.Println(res)
// 		j := cut(res)
// 		fmt.Println(j)
// 	}
// }

// func ubr(str string) string {
// 	if str[0] == ' ' {
// 		return ubr(str[1:])
// 	}

// 	n := len(str)
// 	if str[n-1] == ' ' {
// 		return ubr(str[:n-1])
// 	}
// 	res := ""
// 	for i := 0; i < n; i++ {
// 		if str[i] == ' ' && str[i+1] == ' ' {
// 			continue
// 		}
// 		res += string(str[i])
// 	}
// 	return res
// }

// func cut(str string) string {
// 	new := ""
// 	for i := 0; i < len(str); i++ {
// 		if str[i] == ' ' {

// 			new = str[:i]
// 			str = str[i+1:]
// 			break
// 		}
// 	}
// 	new = str + " " + new
// 	return new
// }

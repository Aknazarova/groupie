package main

import (
	"fmt"
	"os"
	"strconv"
)

const max = 9223372036854775807
const min = -9223372036854775808

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		fmt.Println(0)
		return
	}

	a, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(0)
		return
	}

	b, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println(0)
		return
	}

	sign := args[1]

	switch sign {
	case "+":
		// +a + +b
		if a > 0 && b > 0 && a > max-b {
			fmt.Println(0)
			return
		}

		// -a + -b
		if a > 0 && b < 0 && a < min-b {
			fmt.Println(0)
			return
		}

		fmt.Println(a + b)
	case "-":
		// -a - +b
		if a < 0 && b > 0 && a < min+b {
			fmt.Println(0)
			return
		}

		// +a - -b
		if a > 0 && b < 0 && a > max+b {
			fmt.Println(0)
			return
		}

		fmt.Println(a - b)
	case "*":
		// +a * -b
		if a > 0 && b < 0 && b < min/a {
			fmt.Println(0)
			return
		}

		// -a * +b
		if a < 0 && b > 0 && a < min/b {
			fmt.Println(0)
			return
		}

		// +a * +b
		if a > 0 && b > 0 && a > max/b {
			fmt.Println(0)
			return
		}

		// -a * -b
		if a < 0 && b < 0 && a < max/b {
			fmt.Println(0)
			return
		}

		fmt.Println(a * b)
	case "/":
		// a / b
		if b == 0 {
			fmt.Println("No division by 0")
			return
		}

		fmt.Println(a / b)
	case "%":
		// a % b
		if b == 0 {
			fmt.Println("No modulo by 0")
			return
		}

		fmt.Println(a % b)
	default:
		fmt.Println(0)
	}
}

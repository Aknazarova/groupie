package main

import (
	"fmt"
	"os"
	"strconv"
)

// PrintChessBoard q
func PrintChessBoard(x, y int) {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			if i%2 == 0 && j%2 == 0 || i%2 == 1 && j%2 == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}

		fmt.Println()
	}
}

func main() {
	arg := os.Args[1:]

	if len(arg) != 2 {
		fmt.Println("Error")
	} else {
		a, _ := strconv.Atoi(arg[0])
		b, _ := strconv.Atoi(arg[1])
		PrintChessBoard(a, b)
	}
}

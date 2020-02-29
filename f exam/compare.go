package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	
	if len(arg) != 2 {
		fmt.Println()
	} else {
	
			if arg[0]<arg[1]{
				z01.PrintRune('-')
				z01.PrintRune('1')
			}else if arg[0]==arg[1]{
				z01.PrintRune('0')
			}else if arg[0]>arg[1]{
				z01.PrintRune('1')
			}
		
		z01.PrintRune(10)
	}

}

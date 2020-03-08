package main

import (
	"fmt"
	"os"
)

func main(){
	arg := os.Args[1:]
	var res string
	if len(arg) != 1{
		fmt.Println()
	}else{
		for _, c := range arg[0]{
			if c >= 97 && c <=110 || c >=65 && c <=78 {
				res += string(c + 13)
			}else if c >= 111 && c <=122 || c >= 79 && c <=90 {
				res += string(c - 13)
			}else if c >= 32 && c <=64  {
				res += string(c)
			}
		}
		fmt.Print(res)
	}
	
}
package main

import "fmt"

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func main() {
	num1 := &NodeAddL{Num: 5}
	num1 = pushBack(num1, 1)
	num1 = pushBack(num1, 3)
	num1 = pushBack(num1, 1)
	num1 = pushBack(num1, 3)

	result := Sortll(num1)
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
func pushBack(l *NodeAddL, data int) *NodeAddL {
	n := &NodeAddL{Num: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}
func Sortll(node *NodeAddL) *NodeAddL {
	a := node
	for a != nil {
		b := a.Next
		for b != nil {
			if a.Num < b.Num {
				a.Num, b.Num = b.Num, a.Num
			}
			b = b.Next
		}
		a = a.Next
	}
	return node
}

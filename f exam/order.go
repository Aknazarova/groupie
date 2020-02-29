package main

import "fmt"

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func main() {
	num1 := &NodeAddL{Num: 1}
	num1 = pushBack(num1, 2)
	num1 = pushBack(num1, 3)
	num1 = pushBack(num1, 4)
	num1 = pushBack(num1, 5)

	result := Changeorder(num1)
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
func pushBack(l *NodeAddL, num int) *NodeAddL {
	n := &NodeAddL{Num: num}

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
func Changeorder(node *NodeAddL) *NodeAddL {

	var copy *NodeAddL
	it := node

	for it != nil {
		if it.Num%2 == 1 {
			copy = pushBack(copy, it.Num)
		}
		it = it.Next
	}
	it = node

	for it != nil {
		if it.Num%2 == 0 {
			copy = pushBack(copy, it.Num)
		}
		it = it.Next
	}
	return copy

}

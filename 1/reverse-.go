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
func Reverse(node *NodeAddL) *NodeAddL {

	var prev *NodeL

	for node.Head != nil {
		n := &NodeAddL{}
		n.Next = node.Head.Next
		node.Head.Next = prev
		prev = node.Head
		node.Head = n.Next
	}
	node.Head = prev

}

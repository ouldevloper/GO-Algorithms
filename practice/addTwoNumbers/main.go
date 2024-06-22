package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	remain := 0
	_l1 := l1
	for _l1 != nil && l2 != nil {
		val := (_l1.Val + l2.Val + remain)
		_l1.Val = val % 10
		remain = int(val / 10)
		if _l1.Next == nil { //&& remain != 0 {
			_l1 = &ListNode{10, nil}
		} else {
			_l1 = _l1.Next
		}
		l2 = l2.Next
	}
	fmt.Println(remain)

	return l1
}

func print(l1 *ListNode) {
	for l1 != nil {
		fmt.Println("Item : ", l1.Val)
		l1 = l1.Next
	}
}

func main() {

	l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{8, nil}}}
	l1 = addTwoNumbers(l1, l2)

	print(l1)

}

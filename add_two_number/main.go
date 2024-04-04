package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(l *ListNode) {
	fmt.Println()
	current := l
	for current != nil {
		fmt.Printf("%d -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func add(arr []int) *ListNode {
	head := &ListNode{Val: arr[0]}
	curent := head
	for i := 1; i < len(arr); i++ {
		curent.Next = &ListNode{Val: arr[i]}
		curent = curent.Next
	}
	return head
}

func main() {
	l1, l2 := add([]int{9, 9, 9, 9, 9, 9, 9}), add([]int{9, 9, 9, 9})
	printList(addTwoNumbers(l1, l2))

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	sum := l1.Val + l2.Val
	result := &ListNode{Val: sum % 10}
	c1, c2 := l1.Next, l2.Next
	current := result
	for c1 != nil || c2 != nil {
		if c1 != nil && c2 != nil {
			sum = sum/10 + c1.Val + c2.Val
			c1 = c1.Next
			c2 = c2.Next
		} else if c1 == nil {
			sum = sum/10 + c2.Val
			c2 = c2.Next
		} else if c2 == nil {
			sum = sum/10 + c1.Val
			c1 = c1.Next
		}
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}
	if sum >= 10 {
		current.Next = &ListNode{Val: 1}
	}
	return result
}

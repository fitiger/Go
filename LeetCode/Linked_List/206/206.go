package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// 用于表示 NULL
var null = -1 << 40

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var prev *ListNode
	for head != nil {
		temp := head.Next

		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}
func createList(array []int) *ListNode {
	n := len(array)
	if n == 0 {
		return nil
	}

	head := &ListNode{Val: array[0]}

	node := head
	for i := 1; i < n; i++ {
		node.Next = &ListNode{Val: array[i]}
		node = node.Next
	}
	return head
}

func printList(head *ListNode) {
	for node := head; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}
func main() {
	inputs := []int{1, 1, 1, 2, 3}

	head := createList(inputs)
	printList(head)
	fmt.Println("-------------")
	ans := reverseList(head)
	printList(ans)
}

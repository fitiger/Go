package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var pre *ListNode
	for node := slow; node != nil; {
		pre, node, node.Next = node, node.Next, pre
	}

	first := head
	for pre.Next != nil {
		first.Next, first = pre, first.Next
		pre.Next, pre = first, pre.Next
	}

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
	inputs := []int{1, 2, 3, 4}

	head := createList(inputs)
	printList(head)
	fmt.Println("-------------")
	reorderList(head)
	printList(head)
}

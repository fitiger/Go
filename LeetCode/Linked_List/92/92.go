package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	dummy := &ListNode{0, head}
	head = dummy

	for i := 0; i < left-1; i++ {
		head = head.Next
	}
	var node, prev *ListNode = head.Next, nil

	for i := 0; i < right-left+1; i++ {
		next := node.Next
		node.Next = prev
		prev = node
		node = next
	}
	head.Next.Next = node
	head.Next = prev

	return dummy.Next
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
	inputs := []int{1, 2}

	head := createList(inputs)
	printList(head)
	fmt.Println("-------------")
	ans := reverseBetween(head, 1, 2)
	printList(ans)
}

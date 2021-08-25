package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Bubble sort
func sortList(head *ListNode) *ListNode {
	for i := head; i != nil; i = i.Next {
		for j := i.Next; j != nil; j = j.Next {
			if j.Val < i.Val {
				i.Val, j.Val = j.Val, i.Val
			}
		}
	}
	return head
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
	inputs := []int{1, 2, 4}

	head := createList(inputs)
	printList(head)
	fmt.Println("-------------")
	ans := sortList(head)
	printList(ans)
}

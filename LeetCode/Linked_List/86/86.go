package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	newList := &ListNode{}
	newNode := newList
	dummy := &ListNode{Next: head}
	head = dummy
	for head.Next != nil {
		if head.Next.Val < x {
			head = head.Next
		} else {
			temp := head.Next
			head.Next = head.Next.Next
			newNode.Next = temp
			newNode = newNode.Next
		}
	}
	newNode.Next = nil
	head.Next = newList.Next
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
	inputs := []int{1, 2, 4, 3, 1, 5, 2}

	head := createList(inputs)

	printList(head)
	fmt.Println("-------------")
	ans := partition(head, 3)
	printList(ans)
}

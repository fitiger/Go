package main

import "fmt"

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

//实在缺乏逻辑，需要更简单的逻辑！
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	newList := &ListNode{Next: l1}
	start := newList
	for l2 != nil && l1 != nil {
		if l1.Val < l2.Val {
			newList.Next = l1
			newList = newList.Next
			l1 = l1.Next
		} else {
			newList.Next = l2
			newList = newList.Next
			l2 = l2.Next
		}
	}
	for l1 != nil {
		newList.Next = l1
		newList = newList.Next
		l1 = l1.Next
	}
	for l2 != nil {
		newList.Next = l2
		newList = newList.Next
		l2 = l2.Next
	}
	return start.Next
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
	head2 := createList(inputs)
	printList(head)
	fmt.Println("-------------")
	ans := mergeTwoLists(head, head2)
	printList(ans)
}

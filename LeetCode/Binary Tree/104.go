package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 用于表示 NULL
var NINF = -1 << 63

// Create a new binary tree according to given array.
func CreateTree(array []int) *TreeNode {
	n := len(array)
	if n == 0 {
		return nil
	}

	root := &TreeNode{Val: array[0]}

	queue := make([]*TreeNode, 0, n)
	queue = append(queue, root)

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && array[i] != NINF {
			node.Left = &TreeNode{Val: array[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < n && array[i] != NINF {
			node.Right = &TreeNode{Val: array[i]}
			queue = append(queue, node.Right)
		}
		i++

	}
	return root
}

// 按照输入顺序打印二叉树
func PrintTree(root *TreeNode) {

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		root = queue[0]
		fmt.Println(root.Val)
		if root.Left != nil {
			queue = append(queue, root.Left)
		}
		if root.Right != nil {
			queue = append(queue, root.Right)
		}
		queue = queue[1:]
	}
}

// Find the maximum depth of binary tree
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left) + 1
	right := maxDepth(root.Right) + 1
	if left < right {
		return right
	} else {
		return left
	}
}

func main() {
	inputs := []int{1, 2, 3, 4, 5}

	root := CreateTree(inputs)
	PrintTree(root)
	ans := maxDepth(root)
	fmt.Printf("maximum depth is %d\n", ans)
}

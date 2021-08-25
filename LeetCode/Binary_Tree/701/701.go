package main

import "fmt"

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 用于表示 NULL
var null = -1 << 40

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

		if i < n && array[i] != null {
			node.Left = &TreeNode{Val: array[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < n && array[i] != null {
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

//recursive
func insertIntoBST_recursive(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

//iterative
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	node := root

	for node != nil {

		if node.Val < val {
			if node.Right != nil {
				node = node.Right
			} else {
				node.Right = &TreeNode{Val: val}
				break
			}
		} else {
			if node.Left != nil {
				node = node.Left
			} else {
				node.Left = &TreeNode{Val: val}
				break
			}

		}
	}
	return root
}

func main() {
	inputs := []int{4, 2, 7, 1, 3}

	root := CreateTree(inputs)
	//PrintTree(root)
	ans := insertIntoBST(root, 5)
	PrintTree(ans)
	//fmt.Println(ans)
}

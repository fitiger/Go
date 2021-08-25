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

//递归的时候同时记录深度和是否符合条件
func isBalancedin(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	flag, left := isBalancedin(root.Left)
	if !flag {
		return false, 0
	}

	flag, right := isBalancedin(root.Right)
	if !flag {
		return false, 0
	}

	if left-right > 1 || right-left > 1 {
		return false, 0
	}
	//平衡的话return最大数字
	if left < right {
		return true, right + 1
	} else {
		return true, left + 1
	}

}

//受到题目限制必须要多创建一个函数来完成sad
func isBalanced(root *TreeNode) bool {
	ans, _ := isBalancedin(root)
	return ans
}

func main() {
	inputs := []int{1, 2, 3, 4, 5}

	root := CreateTree(inputs)
	PrintTree(root)
	ans := isBalanced(root)
	fmt.Printf("The ans is %t\n", ans)
}

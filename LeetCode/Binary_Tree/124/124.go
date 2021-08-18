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

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func maxSum(root *TreeNode) (int, int) {

	if root == nil {
		return null, null
	}
	max_total_l, max_left := maxSum(root.Left)
	max_total_r, max_right := maxSum(root.Right)
	max_total := max(max_total_l, max_total_r) //所以子节点路径最大和

	max_solo := max(max(max_left, max_right)+root.Val, root.Val) //单一子节点至自身路径最大和

	max_total = max(max(max_total, max_solo), max_left+max_right+root.Val)
	return max_total, max_solo
}

func maxPathSum(root *TreeNode) int {
	ans, _ := maxSum(root)
	return ans
}

func main() {
	inputs := []int{1, 2, 3}

	root := CreateTree(inputs)
	PrintTree(root)
	ans, _ := maxSum(root)
	fmt.Printf("The ans is %d\n", ans)
}

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

func lcmPath(root *TreeNode, x *TreeNode) (bool, []int) {
	if root == nil {
		return false, make([]int, 0)
	}
	if root.Val == x.Val {
		return true, []int{root.Val}
	}
	flag, path := lcmPath(root.Left, x)
	if flag {
		return true, append(path, root.Val)
	}
	flag, path = lcmPath(root.Right, x)
	if flag {
		return true, append(path, root.Val)
	}
	return false, make([]int, 0)
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == q.Val || root.Val == p.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right
}

func main() {
	inputs := []int{1, 2, 3}
	x := CreateTree([]int{2})
	y := &TreeNode{Val: 3}
	fmt.Println(x, y)
	fmt.Println(x == y)
	root := CreateTree(inputs)
	PrintTree(root)
	ans := lowestCommonAncestor(root, x, y)
	fmt.Println(ans.Val)
	//fmt.Printf("The ans is %d\n", ans)
	//method 2 To do

	_, path_x := lcmPath(root, x)
	fmt.Println(len(path_x))
}

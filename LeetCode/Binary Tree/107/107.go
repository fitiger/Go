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

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	ans := make([][]int, 0)
	ans = append(ans, []int{})
	queue := make([]*TreeNode, 1)
	queue[0] = root
	i := 0
	depth := make([]int, 1000)
	depth[0] = 1
	// next depth
	j := 1
	for i < len(queue) {
		node := queue[i]
		if i == depth[j-1] {
			ans = append(ans, []int{})
			depth[j] += depth[j-1]
			j++
		}
		ans[j-1] = append(ans[j-1], node.Val)
		if node.Left != nil {
			queue = append(queue, node.Left)
			depth[j]++
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
			depth[j]++
		}
		i++
	}
	return ans
}

func levelOrder_notme(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ { //level - to track where we need to insert values
		res = append(res, []int{}) //adding slice for the new "level" of values
		for levelSize := len(queue); levelSize > 0; levelSize-- {
			//levelSize - to track how many elements we need to dequeue and insert in the current "level"
			if queue[0].Left != nil {
				queue = append(queue, queue[0].Left)
			} //adding next nodes to the queue
			if queue[0].Right != nil {
				queue = append(queue, queue[0].Right)
			}
			res[level] = append(res[level], queue[0].Val) //adding first element in the queue to a "level" slice
			queue = queue[1:]                             //deque first element
		}
	}
	return res
}
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := levelOrder_notme(root)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func main() {
	inputs := []int{3, 9, 20, null, null, 15, 7}

	root := CreateTree(inputs)
	PrintTree(root)
	ans := levelOrderBottom(root)
	fmt.Println(ans)
}

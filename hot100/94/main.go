package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	var recur func(*TreeNode)
	recur = func(node *TreeNode) {
		if node == nil {
			return
		}
		recur(node.Left)
		res = append(res, node.Val)
		recur(node.Right)
	}
	recur(root)
	return res
}

func inorderTraversal2(root *TreeNode) (res []int) {
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 中序遍历
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

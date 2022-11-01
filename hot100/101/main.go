package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var recur func(*TreeNode, *TreeNode) bool
	recur = func(left *TreeNode, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		return left.Val == right.Val && recur(left.Left, right.Right) && recur(left.Right, right.Left)
	}
	return recur(root.Left, root.Right)
}

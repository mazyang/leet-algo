package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：中序遍历，遍历的当前节点是否大于保存的上一个的节点历
func isValidBST(root *TreeNode) bool {
	// 比较上一个节点和当前节点是否是递增的
	pre := math.MinInt
	var inorder func(node *TreeNode) bool
	inorder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		left := inorder(node.Left)
		if node.Val <= pre {
			return false
		}
		pre = node.Val
		right := inorder(node.Right)
		return left && right
	}
	return inorder(root)
}

// 方法二：递归，每棵树设置最大值和最小值，判断当前节点的值是否在最大值和最小值之间
func isValidBST2(root *TreeNode) bool {
	var recur func(*TreeNode, int, int) bool
	recur = func(node *TreeNode, min int, max int) bool {
		if node == nil {
			return true
		}
		if node.Val >= max && node.Val <= min {
			return false
		}
		return recur(node.Left, min, node.Val) && recur(node.Right, node.Val, max)
	}
	return recur(root, math.MinInt, math.MaxInt)
}

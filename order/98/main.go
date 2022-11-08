package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：递归
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

// 每棵树设置最大值和最小值，判断当前节点的值是否在最大值和最小值之间
func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}
	// 判断当前节点的值是否在 lower 和 upper 之间
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

// 方法二：中序遍历, 判断遍历的当前节点是否大于保存的上一个的节点
func isValidBST2(root *TreeNode) bool {
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

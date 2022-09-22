package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的直径
// 思路：每⼀条⼆叉树的「直径」⻓度，就是⼀个节点的左右⼦树的最⼤深度之和
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var maxDepth func(node *TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := maxDepth(node.Left)
		rightDepth := maxDepth(node.Right)
		// 后序位置，计算最大直径
		maxDiameter = int(math.Max(float64(leftDepth+rightDepth), float64(maxDiameter)))
		return 1 + int(math.Max(float64(leftDepth), float64(rightDepth)))
	}
	maxDepth(root)
	return maxDiameter
}

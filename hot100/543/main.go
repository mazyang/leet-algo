package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var res int
	// 获得每个节点最大深度的同时计算直径
	var maxDepth func(*TreeNode) int
	maxDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := maxDepth(node.Left)
		rightDepth := maxDepth(node.Right)
		res = max(leftDepth+rightDepth, res)
		return max(leftDepth, rightDepth) + 1
	}
	maxDepth(root)
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 方法一：递归，后序遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return 1 + int(math.Max(float64(left), float64(right)))
}

// 方法二：遍历每个节点，获得所在深度，遇到叶子节点代表有可能是最深的位置
func maxDepth2(root *TreeNode) int {
	ans := 0
	// 使用depth变量记录子节点的深度, 保存到ans中
	depth := 0
	var recur func(*TreeNode)
	recur = func(node *TreeNode) {
		if node == nil {
			return
		}
		// 进入节点
		depth++
		// 说明当前节点是叶子节点，记录此时的深度
		if node.Left == nil && node.Right == nil {
			ans = int(math.Max(float64(ans), float64(depth)))
		}
		recur(node.Left)
		recur(node.Right)
		// 离开节点
		depth--
	}
	recur(root)
	return ans
}

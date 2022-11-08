package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的最大路径和
func maxPathSum(root *TreeNode) int {
	// 存放最大路径和
	res := math.MinInt
	if root == nil {
		return 0
	}
	var recur func(*TreeNode) int
	// 该函数计算二叉树中的一个节点的最大贡献值，具体而言，就是在以该节点为根节点的子树中寻找以该节点为起点的一条路径，使得该路径上的节点值之和最大
	recur = func(node *TreeNode) int {
		// 空节点的最大贡献值为 0
		if node == nil {
			return 0
		}
		// 计算当前节点的左节点的最大贡献值，如果小于0，则不计入
		leftMaxSum := max(0, recur(node.Left))
		// 计算当前节点的右节点的最大贡献值，如果小于0，则不计入
		rightMaxSum := max(0, recur(node.Right))

		// 下面两行代码计算最大路径和
		// 当前节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
		pathMaxSum := node.Val + leftMaxSum + rightMaxSum
		// 当前的最大路径和 与 res 存放的最大路径和比较，选择最大的一个
		res = max(res, pathMaxSum)

		// 非空节点的最大贡献值等于节点值与其子节点中的最大贡献值之和
		return max(leftMaxSum, rightMaxSum) + node.Val
	}
	recur(root)
	return res
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

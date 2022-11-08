package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 判断是否为平衡二叉树
// 方法一：自底向上，计算子树的高度，并且判断是否平衡
func isBalanced(root *TreeNode) bool {
	var recur func(*TreeNode) int
	recur = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := recur(node.Left)
		// 为什么不统一判断：left==-1 || right==-1？因为可以减少重复计算
		if left == -1 {
			return -1
		}
		right := recur(node.Right)
		if right == -1 {
			return -1
		}
		// 如果当前子树是平衡二叉树，就返回树的高度，否则返回-1
		if math.Abs(float64(left-right)) < 2 {
			return int(math.Max(float64(left), float64(right)) + 1)
		} else {
			return -1
		}
	}
	return recur(root) != -1
}

// 方法二：自顶向下暴力求解
func isBalanced2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 先计算根节点是否平衡，然后计算依次左右子树是否平衡
	return math.Abs(float64(height(root.Left)-height(root.Right))) <= 1 && isBalanced2(root.Left) && isBalanced2(root.Right)
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(height(root.Left)), float64(height(root.Right))) + 1)
}

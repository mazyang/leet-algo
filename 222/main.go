package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	left := root.Left
	right := root.Right
	hl, hr := 0, 0
	for left != nil {
		left = left.Left
		hl++
	}
	for right != nil {
		right = right.Right
		hr++
	}
	// 说明是一颗满二叉树
	if hl == hr {
		return int(math.Pow(2, float64(hl)) - 1)
	}
	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var build func(int, int) *TreeNode
	build = func(left int, right int) *TreeNode {
		if left > right {
			return nil
		}
		// 寻找最大值及其索引
		index := -1
		maxVal := math.MinInt
		for i := left; i <= right; i++ {
			if maxVal < nums[i] {
				maxVal = nums[i]
				index = i
			}
		}
		root := &TreeNode{Val: maxVal}
		l := build(left, index-1)
		r := build(index+1, right)
		root.Left = l
		root.Right = r
		return root
	}
	return build(0, len(nums)-1)
}

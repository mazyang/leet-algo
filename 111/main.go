package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的最小深度
// 方法一：层次遍历，找到一个节点的左右都为空，保存层次
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
	queue := []*TreeNode{root}
	res := 0
	for len(queue) != 0 {
		depth++
		for i := len(queue); i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return res
}

// 方法二：深度优先搜索：分别计算其左右子树的最小叶子节点深度
func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	m1 := minDepth2(root.Left)
	m2 := minDepth2(root.Right)
	// 如果左右子树有一个为空, 返回不为空的子树的深度, 都为空则返回 1
	if root.Left == nil || root.Right == nil {
		return m1 + m2 + 1
	}
	// 如果左右子树都不为空
	return int(math.Min(float64(m1), float64(m2)) + 1)
}

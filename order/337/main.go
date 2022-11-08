package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var dp func(*TreeNode) []int
	// 返回值 arr[0] 代表不抢node节点获得的最大收益，arr[1]代表抢node节点获得的最大收益
	dp = func(node *TreeNode) []int {
		if node == nil {
			return []int{0, 0}
		}
		left := dp(node.Left)
		right := dp(node.Right)
		// 抢当前节点，左右节点不抢
		do := node.Val + left[0] + right[0]
		// 不抢当前节点，左节点可以抢也可以不抢(取最大的)，右节点同样
		donot := max(left[0], left[1]) + max(right[0], right[1])
		return []int{donot, do}
	}
	res := dp(root)
	return max(res[0], res[1])
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob2(root *TreeNode) int {
	memo := make(map[*TreeNode]int)
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if _, ok := memo[root]; ok {
			return memo[root]
		}
		a := dfs(root.Left) + dfs(root.Right)
		b := root.Val
		if root.Left != nil {
			b += dfs(root.Left.Left) + dfs(root.Left.Right)
		}
		if root.Right != nil {
			b += dfs(root.Right.Left) + dfs(root.Right.Right)
		}
		res := max(a, b)
		memo[root] = res
		return res
	}
	return dfs(root)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

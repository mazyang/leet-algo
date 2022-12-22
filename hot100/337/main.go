package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 打家劫舍3
func rob(root *TreeNode) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	// 返回一个大小为 2 的数组 arr
	//arr[0] 表示不抢 root 的话，得到的最大钱数
	//arr[1] 表示抢 root 的话，得到的最大钱数
	var dfs func(*TreeNode) []int
	dfs = func(node *TreeNode) []int {
		if node == nil {
			return []int{0, 0}
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		// 不抢当前节点
		not_rob := max(left[0], left[1]) + max(right[0], right[1])
		// 抢当前节点
		rob := node.Val + left[0] + right[0]
		return []int{not_rob, rob}
	}
	res := dfs(root)
	return max(res[0], res[1])
}

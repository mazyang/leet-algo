package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	var ans int
	// preSum[i]=j 代表从二叉树的根节点开始，路径和为i的路径有j条
	preSum := map[int64]int{0: 1}
	var dfs func(*TreeNode, int64)
	// curr 代表从root节点到当前节点的前缀和
	dfs = func(node *TreeNode, curr int64) {
		if node == nil {
			return
		}
		curr += int64(node.Val)
		ans += preSum[curr-int64(targetSum)]
		// 路径和为curr的个数加 1
		preSum[curr]++
		dfs(node.Left, curr)
		dfs(node.Right, curr)
		preSum[curr]--
		return
	}
	dfs(root, 0)
	return ans
}

package main

// 不同路径
// 方法：动态规划
func uniquePaths(m int, n int) int {
	// 状态是机器人所在位置，选择只有向下或向右
	// dp[i][j]表示代表以位置 (i, j) 为结尾的路径个数
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	// base case
	// 初始化第一行和第一列
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

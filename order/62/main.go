package main

func uniquePaths(m int, n int) int {
	// 状态是机器人所在位置，选择只有向下或向右
	// dp[i][j] 代表以 i,j 作为终点的路径
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	// base case
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

func main() {
	println(uniquePaths(3, 2))
}

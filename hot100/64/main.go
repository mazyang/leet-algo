package main

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[0][0] = grid[0][0]
				continue
			}
			if i == 0 {
				dp[i][j] = grid[i][j] + dp[i][j-1]
			} else if j == 0 {
				dp[i][j] = grid[i][j] + dp[i-1][j]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

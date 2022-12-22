package main

// 最大正方形
func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	// dp[i][j]代表以 matrix[i-1][j-1] 作为右下角的最大正方形的变长
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	maxSlide := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i+1][j+1] = min(min(dp[i][j+1], dp[i+1][j]), dp[i][j]) + 1
				if dp[i+1][j+1] > maxSlide {
					maxSlide = dp[i+1][j+1]
				}
			}
		}
	}
	return maxSlide * maxSlide
}

package main

import (
	"fmt"
	"math"
)

// 方法一：暴力回溯
func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	res := 20000
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	var dfs func(int, int) int
	dfs = func(row int, col int) int {
		if row < 0 || col < 0 || col >= n || row >= n {
			return 20000
		}
		if row == 0 {
			return matrix[row][col]
		}
		return matrix[row][col] + min(dfs(row-1, col), min(dfs(row-1, col-1), dfs(row-1, col+1)))
	}
	// 终点可能在最后一行的任意一列
	for i := 0; i < n; i++ {
		res = min(res, dfs(n-1, i))
	}
	return res
}

// 方法二：递归+备忘录
func minFallingPathSum2(matrix [][]int) int {
	n := len(matrix)
	// memo[i][j] 记录在i,j位置的路径最小和
	memo := make([][]int, n)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = 66666
		}
	}
	res := math.MaxInt
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	var dfs func(int, int) int
	dfs = func(row int, col int) int {
		if row < 0 || col < 0 || col >= n || row >= n {
			return 66666
		}
		if row == 0 {
			return matrix[row][col]
		}
		if memo[row][col] != 66666 {
			return memo[row][col]
		}
		memo[row][col] = matrix[row][col] + min(dfs(row-1, col), min(dfs(row-1, col-1), dfs(row-1, col+1)))
		return memo[row][col]
	}
	// 终点可能在最后一行的任意一列
	for i := 0; i < n; i++ {
		res = min(res, dfs(n-1, i))
	}
	return res
}

// 方法三：动态规划, 自底向上
func minFallingPathSum3(matrix [][]int) int {
	n := len(matrix)
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	// dp[i][j] 记录在i,j位置的路径最小和
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	// 初始化第一行
	for i := 0; i < n; i++ {
		dp[0][i] = matrix[0][i]
	}
	// 行状态
	for i := 1; i < n; i++ {
		// 列状态
		for j := 0; j < n; j++ {
			// 对于j=0和j=n-1 特殊处理
			if j == 0 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
			} else if j == n-1 {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + matrix[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], min(dp[i-1][j-1], dp[i-1][j+1])) + matrix[i][j]
			}
		}
	}
	res := math.MaxInt
	for i := 0; i < n; i++ {
		res = min(res, dp[n-1][i])
	}
	return res
}

func main() {
	matrix := [][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}
	fmt.Println(minFallingPathSum3(matrix))
}

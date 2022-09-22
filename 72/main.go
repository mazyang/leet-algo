package main

import "fmt"

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 方法一：暴力递归
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	// dfs(i, j) 的返回值就是 word1[0...i] word2[0...j] 的最小编辑距离
	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		// 进行状态转移
		if word1[i] == word2[j] {
			return dfs(i-1, j-1)
		} else {
			return min(
				dfs(i, j-1), // 插入
				min(dfs(i-1, j), // 删除
					dfs(i-1, j-1))) + 1 // 替换
		}
	}
	// i j 初始化指向最后一个索引
	return dfs(m-1, n-1)
}

// 方法二：暴力递归+备忘录
func minDistance2(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	memo := make([][]int, m)
	for i := 0; i < m; i++ {
		memo[i] = make([]int, n)
	}
	// dfs(i, j) 的返回值就是 word1[0...i] word2[0...j] 的最小编辑距离
	var dfs func(int, int) int
	dfs = func(i int, j int) int {
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		if memo[i][j] != 0 {
			return memo[i][j]
		}
		// 进行状态转移
		if word1[i] == word2[j] {
			memo[i][j] = dfs(i-1, j-1)
		} else {
			memo[i][j] = min(
				dfs(i, j-1), // 插入：在i的位置插入一个和word2[j]相同的字符，j向前移动，继续和位置i上的元素对比
				min(dfs(i-1, j), // 删除：直接把word1[i]删除，i向前移动，继续和j位置的字符比较
					dfs(i-1, j-1))) + 1 // 替换：将word1[i]替换成word2[j]，i,j同时向前移动
		}
		return memo[i][j]
	}
	// i j 初始化指向最后一个索引
	return dfs(m-1, n-1)
}

// 方法三：动态规划
func minDistance3(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	// base case
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
			}
		}
	}
	return dp[m][n]
}

func main() {
	word1 := "horse"
	word2 := "ros"
	fmt.Println(minDistance(word1, word2))
}

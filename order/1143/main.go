package main

import "fmt"

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 方法一：暴力递归
func longestCommonSubsequence(text1 string, text2 string) int {
	var dfs func(int, int) int
	// text1=babcde   text2=acze
	dfs = func(i int, j int) int {
		if i == -1 || j == -1 {
			return 0
		}
		// 找到一个 最长公共子序列 中的元素
		if text1[i] == text2[j] {
			return dfs(i-1, j-1) + 1
		} else {
			return max(dfs(i-1, j), dfs(i, j-1))
		}
	}
	return dfs(len(text1)-1, len(text2)-1)
}

// 方法二：动态规划
func longestCommonSubsequence2(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
	// dp[i][j] 的含义：对于 text1[0:i-1] 和 text2[0:j-1] 它们的 LCS 的长度为 dp[i][j]
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			}
		}
	}
	return dp[m][n]
}

func main() {
	str1 := "babcde"
	str2 := "acze"
	fmt.Println(longestCommonSubsequence(str1, str2))
}

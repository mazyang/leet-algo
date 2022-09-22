package main

// 动态规划
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i][j-1], dp[i+1][j])
			}
		}
	}
	return dp[0][n-1]
}

// 状态压缩
func longestPalindromeSubseq2(s string) int {
	n := len(s)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := n - 2; i >= 0; i-- {
		pre := 0
		for j := i + 1; j < n; j++ {
			tmp := dp[j]
			if s[i] == s[j] {
				dp[j] = pre + 2
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			pre = tmp
		}
	}
	return dp[n-1]
}

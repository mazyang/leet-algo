package main

// 方法一：动态规划
func countSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	var res int
	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			if s[i] == s[j] && (j-i < 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				res++
			}
		}
	}
	return res
}

// 方法二：中心扩展法
func countSubstrings2(s string) int {
	res := 0
	// 共有 2n-1个中心点，以单个字符为中心点的个数为n，以两个字符为中心点的个数为n-1
	for center := 0; center < 2*len(s)-1; center++ {
		// (left, right)=(0 0) (0 1) (1 1) (1 2) (2 2) (2 3) ...
		left := center / 2
		right := left + center%2
		for left >= 0 && right < len(s) && s[left] == s[right] {
			res++
			left--
			right++
		}
	}
	return res
}

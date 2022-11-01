package main

// 方法一：遍历字符串，每次以当前字符为中心或者以当前字符和下一个字符为中心向外扩散，得到最长的回文串
func longestPalindrome(s string) string {
	var res string
	// 从中间向外扩展
	for i := 0; i < len(s); i++ {
		s1 := subString(i, i, s)
		s2 := subString(i, i+1, s)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func subString(i, j int, s string) string {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	// s[i+1:j] 是左闭右开的区间
	return s[i+1 : j]
}

// 方法二：动态规划
func longestPalindrome2(s string) string {
	n := len(s)
	// dp[i][j]的含义是：s[i:j]是否为回文串
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}
	mx, start := 1, 0
	for j := 0; j < n; j++ {
		for i := 0; i <= j; i++ {
			// s[i:j]中只有一个或两个字符
			if j-i < 2 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
			}
			// 更新最长回文串的起始和结束索引
			if dp[i][j] && mx < j-i+1 {
				mx, start = j-i+1, i
			}
		}
	}
	return s[start : start+mx]
}

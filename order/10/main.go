package main

import "strconv"

// 动态规划
func isMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	// base case：如果s为空，p不为空，也有可能为true
	for j := 1; j <= len(p); j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				if s[i-1] == p[j-2] || p[j-2] == '.' {
					dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j]
				} else {
					// p[j-1] 重复 0个 p[j-2]
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}

// 递归+备忘录
func isMatch2(s string, p string) bool {
	m, n := len(s), len(p)
	memo := make(map[string]bool)
	var dp func(int, int) bool
	dp = func(i int, j int) bool {
		if j == n {
			return i == m
		}
		// 当i走到末尾时，j没有走到末尾
		if i == m {
			if (n-j)%2 == 1 {
				return false
			}
			for ; j+1 < n; j += 2 {
				if p[i+1] != '*' {
					return false
				}
			}
		}
		key := strconv.Itoa(i) + "," + strconv.Itoa(j)
		if _, ok := memo[key]; ok {
			return memo[key]
		}
		var res bool
		if s[i] == p[j] || p[j] == '.' {
			if j < n-1 && p[j+1] == '*' {
				res = dp(i, j+2) || dp(i+1, j)
			} else {
				res = dp(i+1, j+1)
			}
		} else {
			if j < n-1 && p[j+1] == '*' {
				res = dp(i, j+2)
			} else {
				res = false
			}
		}
		memo[key] = res
		return res
	}
	return dp(0, 0)
}

func main() {
	s := "aab"
	p := "c*a*b"
	isMatch(s, p)
}

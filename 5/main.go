package main

import "fmt"

// 最长回文子串
// 方法一：遍历字符串，每次以当前字符为中心或者以当前字符和下一个字符为中心向外扩散，得到最长的回文串
func longestPalindrome(s string) string {
	var res string
	for i := 0; i < len(s); i++ {
		s1 := palindrome(s, i, i)
		s2 := palindrome(s, i, i+1)
		if len(s1) > len(res) {
			res = s1
		}
		if len(s2) > len(res) {
			res = s2
		}
	}
	return res
}

func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

// 方法二：动态规划
func longestPalindrome2(s string) string {
    n := len(s)
    dp := make([][]bool, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]bool, n)
    }
    mx, start := 1, 0
    for j := 0; j < n; j++ {
        for i := 0; i <= j; i++ {
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

func main() {
	s := "abcdef"
	fmt.Println(palindrome(s, 5, 5))
}
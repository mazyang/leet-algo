package main

import "math"

func longestValidParentheses(s string) int {
	// stack 放置左括号的索引
	var stack []int
	// dp[i]记录以s[i-1]为结尾的最长合法括号子串长度
	dp := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			dp[i+1]=0
		} else {
			if len(stack) != 0 {
				leftIndex := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				len := 1 + i - leftIndex + dp[leftIndex]
				dp[i+1] = len
			} else {
				dp[i+1] = 0
			}
		}
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		res = int(math.Max(float64(res), float64(dp[i])))
	}
	return res
}
package main

import "math"

// 最长有效括号
// 方法一：栈+动态规划
func longestValidParentheses(s string) int {
	var stack []int
	// dp[i]记录以s[i-1]为结尾的最长合法括号子串长度
	dp := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			dp[i+1] = 0
		} else {
			if len(stack) != 0 {
				// 最近的左括号的索引
				leftIndex := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				// 计算合法括号的长度
				length := 1 + i - leftIndex + dp[leftIndex]
				dp[i+1] = length
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

func longestValidParentheses2(s string) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	left, right, maxLength := 0, 0, 0
	// 从左到右遍历
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if right > left { // 右括号较多，重置
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if right < left {
			left, right = 0, 0
		}
	}
	return maxLength
}

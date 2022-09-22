package main

import "math"

// 递归+备忘录
func rob(nums []int) int {
	memo := make([]int, len(nums))
	for i := 0; i < len(memo); i++ {
		memo[i] = -1
	}
	var dp func(int) int
	dp = func(start int) int {
		if start >= len(nums) {
			return 0
		}
		if memo[start] != -1 {
			return memo[start]
		}
		res := int(math.Max(float64(dp(start+1)), float64(nums[start]+dp(start+2))))
		memo[start] = res
		return res
	}
	return dp(0)
}

// 动态规划
func rob2(nums []int) int {
	// dp[i] 代表从i开始抢劫获得的钱
	dp := make([]int, len(nums)+2)
	for i := len(nums) - 1; i >= 0; i-- {
		dp[i] = int(math.Max(float64(dp[i+1]), float64(nums[i]+dp[i+2])))
	}
	return dp[0]
}

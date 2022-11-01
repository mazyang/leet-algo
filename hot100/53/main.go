package main

import "math"

// 方法一：动态规划
func maxSubArray(nums []int) int {
	n := len(nums)
	// dp[i] 代表已 nums[i] 为结尾的子数组的和
	dp := make([]int, n)
	dp[0] = nums[0]
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
	}
	res := math.MinInt
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}

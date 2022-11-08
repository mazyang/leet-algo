package main

import "math"

func maxSubArray(nums []int) int {
	n := len(nums)
	// dp[i] 记录以nums[i]为结尾的 最大子数组和
	// dp[i] 有两种「选择」，要么与前⾯的相邻⼦数组连接，形成⼀个和更⼤的⼦数组；要么不与前⾯的⼦数组连接，⾃成⼀派，⾃⼰作为⼀个⼦数组
	dp1, dp2 := nums[0], 0
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	res := nums[0]
	for i := 1; i < n; i++ {
		dp2 = max(nums[i], dp1+nums[i])
		dp1 = dp2
		res = max(res, dp2)
	}
	return res
}

// 方法二：前缀和
func maxSubArray2(nums []int) int {
	n := len(nums)
	preSum := make([]int, n+1)
	preSum[0] = 0
	// preSum[i]保存nums[0] 到 nums[i-1] 的累加和
	for i := 1; i < n+1; i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	res := math.MinInt
	minVal := math.MaxInt // minVal 保存 preSum[0...i] 的最小值
	for i := 0; i < n; i++ {
		minVal = min(minVal, preSum[i])
		res = max(res, preSum[i+1]-minVal)
	}
	return res
}

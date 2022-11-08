package main

func lengthOfLIS(nums []int) int {
	n := len(nums)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	res := 1
	// dp[i]代表以nums[i]结尾的一个最长递增子序列的长度
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

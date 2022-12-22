package main

// 打家劫舍1
func rob(nums []int) int {
	n := len(nums)
	// dp[i] 代表从房间i开始偷取的最大金额，n+2是为了便于计算
	dp := make([]int, n+2)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := n - 1; i >= 0; i-- {
		// 在当前房间可以选择偷或不偷
		dp[i] = max(dp[i+1], dp[i+2]+nums[i])
	}
	return dp[0]
}

func rob2(nums []int) int {
	n := len(nums)
	dp_i := 0
	dp_i_1, dp_i_2 := 0, 0 // 代表dp[i+1] dp[i+2]
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := n - 1; i >= 0; i-- {
		// 从dp[i]开始偷取，如果在当前位置i偷取，下次偷取的位置就从i+2开始，如果当前位置不偷取，那就从i+1开始
		dp_i = max(dp_i_1, nums[i]+dp_i_2)
		dp_i_2 = dp_i_1
		dp_i_1 = dp_i
	}
	return dp_i
}

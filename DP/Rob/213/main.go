package main

// 打家劫舍2
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var recur func(int, int) int
	recur = func(start int, end int) int {
		dp_i := 0
		dp_i_1, dp_i_2 := 0, 0
		for i := end; i >= start; i-- {
			// 在当前房间可以选择偷或不偷
			dp_i = max(dp_i_1, dp_i_2+nums[i])
			dp_i_2, dp_i_1 = dp_i_1, dp_i
		}
		return dp_i
	}
	return max(recur(0, n-2), recur(1, n-1))
}

package main

import "math"

// 动态规划
func maxProduct(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	res := math.MinInt
	// imax 为当前连续的最大值乘积, imin为当前连续的最小值乘积
	imax, imin := 1, 1
	for i := 0; i < len(nums); i++ {
		// 如果当前元素小于0，尝试反转，最大值变为最小值，最小值变为最大值
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		// 乘上当前元素的乘积 和 以当前元素为开头的子数组 判断大小
		imax = max(imax*nums[i], nums[i])
		imin = min(imin*nums[i], nums[i])
		res = max(res, imax)
	}
	return res
}

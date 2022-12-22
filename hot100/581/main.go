package main

import (
	"math"
)

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	minn, maxn := math.MaxInt, math.MinInt
	left, right := -1, -1
	for i := 0; i < n; i++ {
		// 从左到右寻找最大值，寻找右边界
		// 右边界后面的元素都大于maxn
		if maxn > nums[i] {
			right = i
		} else {
			maxn = nums[i]
		}
		// 从右到左寻找最小值，寻找左边界
		// 左边界前面的元素都小于minn
		if minn < nums[n-i-1] {
			left = n - i - 1
		} else {
			minn = nums[n-i-1]
		}
	}
	if right == -1 {
		return 0
	}
	return right - left + 1
}

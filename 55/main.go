package main

func canJump(nums []int) bool {
	n := len(nums)
	// 记录能跳到的最远距离
	farthest := 0
	for i := 0; i < n-1; i++ {
		if farthest < nums[i]+i {
			farthest = nums[i] + i
		}
		// 如果下一步不能跳了, 而且没有到达最后一个元素
		if farthest <= i {
			return false
		}
	}
	return farthest >= n-1
}

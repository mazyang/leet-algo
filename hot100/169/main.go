package main

// 多数元素：使用 Boyer-Moore 投票算法
func majorityElement(nums []int) int {
	candidate := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			count++
		} else if count == 0 {
			candidate = nums[i]
			count = 1
		} else {
			count--
		}
	}
	return candidate
}

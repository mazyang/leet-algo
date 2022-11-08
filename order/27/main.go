package main

// 移除元素
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	fast, slow := 0, 0
	for fast < len(nums) {
		// 如果和val相同，则跳过，不相同就覆盖
		// slow 之前的元素都是与val不同的
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

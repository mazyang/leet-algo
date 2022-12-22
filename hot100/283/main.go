package main

func moveZeroes(nums []int) {
	fast, slow := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			// 覆盖
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
}

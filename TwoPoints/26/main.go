package main

// 删除有序数组中的重复项
func removeDuplicates(nums []int) int {
	slow, fast := 0, 0
	for i := 0; i < len(nums); i++ {
		// 遇到不相同的数字，开始交换
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

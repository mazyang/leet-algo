package main

// 删除有序数组中的重复项, 获得新数组的长度
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	fast, slow := 0, 0
	for fast < len(nums) {
		if nums[fast] != nums[slow] {
			slow++
			nums[slow] = nums[fast]
		}
		fast++
	}
	return slow + 1
}

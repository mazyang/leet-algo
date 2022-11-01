package main

// 搜索旋转排序数组
func search(nums []int, target int) int {
	n := len(nums)
	if n == 1 {
		if target == nums[0] {
			return 0
		} else {
			return -1
		}
	}
	l, r := 0, n-1
	// nums = [4,5,6,7,0,1,2] target=0
	for l <= r {
		mid := (l + r) / 2
		// 找到 target
		if nums[mid] == target {
			return mid
		}
		// 说明左侧是有序的
		if nums[0] <= nums[mid] {
			// 如果 target 在左侧，缩小搜索范围
			if nums[0] <= target && target < nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // 右侧是有序的
			if nums[mid] < target && target <= nums[n-1] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}

package main

import "sort"

// 两数之和, 可能有多对，而且有重复元素
func twoSum(nums []int, target int) [][]int {
	// 排序
	sort.Ints(nums)
	// 双指针移动
	left, right := 0, len(nums)-1
	var res [][]int
	for left < right {
		sum := nums[left] + nums[right]
		leftVal, rightVal := nums[left], nums[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			res = append(res, []int{leftVal, rightVal})
			// 有重复元素，跳过
			for left < right && leftVal == nums[left] {
				left++
			}
			for left < right && rightVal == nums[right] {
				right--
			}
		}
	}
	return res
}

// 三数之和为零
func threeSum(nums []int) [][]int {
	// 先从中取一个数字，然后从剩下的数字中取两个，变为求两数之和
	sort.Ints(nums)
	n := len(nums)
	var res [][]int
	for i := 0; i < n; i++ {
		// 确认第一个数字，之后的两个数字在 nums[i+1:] 中查找
		twoSumList := twoSum(nums[i+1:], 0-nums[i])
		// 取出两数数组，然后加上第一个数字
		for _, val := range twoSumList {
			val = append(val, nums[i])
			res = append(res, val)
		}
		// 遇到第一个数字重复
		for i < n-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}

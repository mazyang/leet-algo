package main

import "math/rand"

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	target := n - k
	left, right := 0, n-1
	for {
		pivotIdx := partition(nums, left, right)
		if pivotIdx == target {
			return nums[pivotIdx]
		} else if pivotIdx > target {
			right = pivotIdx - 1
		} else if pivotIdx < target {
			left = pivotIdx + 1
		}
	}
}

func partition(nums []int, left, right int) int {
	randomIdx := left + rand.Intn(right-left+1)
	nums[randomIdx], nums[left] = nums[left], nums[randomIdx]
	pivot := nums[left]
	i, j := left+1, right
	for i <= j {
		for i <= j && nums[i] <= pivot {
			i++
		}
		for i <= j && nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 找到可以交换的位置j
	nums[left], nums[j] = nums[j], nums[left]
	return j
}

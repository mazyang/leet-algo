package main

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	sort.Sort(sort.Reverse(sort.IntSlice(nums))) // 减少搜索量
	bucket := make([]int, k)
	var dfs func(int) bool
	dfs = func(index int) bool {
		if index == len(nums) {
			return true
		}
		for i := 0; i < k; i++ {
			if bucket[i]+nums[i] > target {
				continue
			}
			if i > 0 && bucket[i] == bucket[i-1] {
				continue
			}
			bucket[i] += nums[index]
			if dfs(index + 1) {
				return true
			}
			bucket[i] -= nums[index]
		}
		return false
	}
	return dfs(0)
}

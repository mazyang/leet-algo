package main

import "sort"

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var path []int
	used := make(map[int]bool)
	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] == true {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack()
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	backtrack()
	return res
}

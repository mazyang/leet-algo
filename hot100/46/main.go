package main

// 全排列
func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	used := make([]bool, len(nums))
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
			path = append(path, nums[i])
			used[i] = true
			backtrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack()
	return res
}

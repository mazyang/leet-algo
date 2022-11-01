package main

func permute(nums []int) [][]int {
	var res [][]int
	var path []int
	used := make([]bool, len(nums))
	var recur func()
	recur = func() {
		if len(nums) == len(path) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] == true {
				continue
			}
			path = append(path, nums[i])
			used[i] = true
			recur()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	recur()
	return res
}

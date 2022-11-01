package main

func combinationSum(candidates []int, target int) [][]int {
	var path []int
	var res [][]int
	var backtrack func(int, int)
	backtrack = func(sum int, start int) {
		if sum == target {
			res = append(res, append([]int{}, path...))
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			backtrack(sum+candidates[i], i)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, 0)
	return res
}

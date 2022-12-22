package main

// 回溯
func findTargetSumWays(nums []int, target int) int {
	var res int
	var dfs func(int, int)
	dfs = func(index int, sum int) {
		if index == len(nums) {
			if sum == target {
				res++
			}
			return
		}
		dfs(index+1, sum+nums[index])
		dfs(index+1, sum-nums[index])
	}
	dfs(0, 0)
	return res
}

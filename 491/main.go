package main

import "fmt"

func findSubsequences(nums []int) [][]int {
	var ans [][]int
	var path []int
	var dfs func(index, last int)
	dfs = func(index, last int) {
		if index == len(nums) {
			if len(path) > 1 {
				cp := make([]int, len(path))
				copy(cp, path)
				ans = append(ans, cp)
			}
			return
		}
		// 满足递增序列
		if nums[index] >= last {
			path = append(path, nums[index])
			dfs(index+1, nums[index])
			path = path[:len(path)-1]
		}
		// 避免重复
		if nums[index] != last {
			dfs(index+1, last)
		}
	}
	dfs(0, -1000)
	return ans
}

// 错误答案，没有解决重复
func findSubsequences2(nums []int) [][]int {
	var path []int
	var res [][]int
	var dfs func(int)
	dfs = func(start int) {
		if len(path) > 1 {
			res = append(res, append([]int{}, path...))
		}
		if start == len(nums) {
			return
		}
		for i := start; i < len(nums)-1; i++ {
			if i > start && nums[i] == nums[i-1] {
				continue
			}
			if nums[i] >= path[len(path)-1] {
				path = append(path, nums[i])
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}
	for i := 0; i < len(nums)-1; i++ {
		path = append(path, nums[i])
		dfs(i + 1)
		path = path[:len(path)-1]
	}
	return res
}

func main() {
	nums := []int{4, 6, 7, 7}
	fmt.Println(findSubsequences(nums))
}

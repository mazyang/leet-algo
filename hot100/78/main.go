package main

import "fmt"

func subsets(nums []int) [][]int {
	var path []int
	var res [][]int
	var backtrack func(int)
	backtrack = func(start int) {
		res = append(res, append([]int{}, path...))
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}

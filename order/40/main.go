package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var path []int
	var res [][]int
	// 先排序
	sort.Ints(candidates)
	var backtrack func(int, int)
	// start 参数是为了避免重复使用元素
	backtrack = func(sum int, start int) {
		if sum == target {
			res = append(res, append([]int{}, path...))
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			// 如果出现重复元素
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			path = append(path, candidates[i])
			backtrack(sum+candidates[i], i+1)
			path = path[:len(path)-1]
		}
	}
	backtrack(0, 0)
	return res
}

func main() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

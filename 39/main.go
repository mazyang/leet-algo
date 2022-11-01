package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int
	var sum int
	var backtrack func(int)
	backtrack = func(start int) {
		if sum == target {
			res = append(res, append([]int{}, path...))
			return
		}
		if sum > target {
			return
		}
		for i := start; i < len(candidates); i++ {
			sum += candidates[i]
			path = append(path, candidates[i])
			// 下一个可以选择当前元素之后的元素（包括当前元素）
			backtrack(i)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}
	backtrack(0)
	return res
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

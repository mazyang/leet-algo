package main

import (
	"fmt"
	"sort"
)

func makesquare(matchsticks []int) bool {
	totalLen := 0
	for _, l := range matchsticks {
		totalLen += l
	}
	if totalLen%4 != 0 {
		return false
	}
	target := totalLen / 4
	sort.Sort(sort.Reverse(sort.IntSlice(matchsticks))) // 减少搜索量

	edges := [4]int{}
	var dfs func(int) bool
	dfs = func(idx int) bool {
		if idx == len(matchsticks) {
			return true
		}
		for i := range edges {
			if i > 0 && edges[i] == edges[i-1] {
				continue
			}
			// 一个火柴在一个edge上可以放入边上，也可以不放入
			edges[i] += matchsticks[idx]
			if edges[i] <= target && dfs(idx+1) {
				return true
			}
			edges[i] -= matchsticks[idx]
		}
		return false
	}
	return dfs(0)
}

func main() {
	//matchsticks := []int{1, 1, 2, 2, 2}
	matchsticks := []int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3}
	fmt.Println(makesquare(matchsticks))
}

package main

// 方法一：回溯 (深度优先搜索)
func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	path := []int{0}
	var res [][]int
	var dfs func(int)
	// 因为是无环图，所以不需要使用visited数组记录是否访问
	dfs = func(index int) {
		if index == n-1 {
			res = append(res, append([]int{}, path...))
			return
		}
		for _, val := range graph[index] {
			path = append(path, val)
			dfs(val)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

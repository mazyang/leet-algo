package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	n := len(graph)
	path := []int{0}
	var res [][]int
	var dfs func(int)
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

func main() {
	graph := [][]int{{1, 2}, {3}, {3}, {}}
	fmt.Println(allPathsSourceTarget(graph))
}

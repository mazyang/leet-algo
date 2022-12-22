package main

import "fmt"

// 所有可能的路径
// 无环图，所有不使用visited数组
func allPathsSourceTarget(graph [][]int) [][]int {
	// graph[0] 存放节点0指向的所有节点
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

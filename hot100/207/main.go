package main

import "fmt"

// 深度优先搜索
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses) // 邻接表
	// visited 记录哪些节点被遍历过，而 onPath 记录当前递归堆栈中有哪些节点
	visited := make([]bool, numCourses) // 节点是否访问过
	onPath := make([]bool, numCourses)  // 记录当前遍历到的路径
	hasCycle := false                   // 判断是否有环
	var dfs func(u int)
	dfs = func(s int) {
		// 如果发现 onPath[s] 已经被标记，说明出现了环
		if onPath[s] {
			hasCycle = true
		}
		if visited[s] || hasCycle {
			// 如果已经被遍历过，或者已经找到了环，那就不用再遍历了，结束递归
			return
		}
		visited[s] = true
		onPath[s] = true
		for _, t := range graph[s] {
			dfs(t)
		}
		onPath[s] = false
	}
	// 建立邻接表
	for _, info := range prerequisites {
		graph[info[1]] = append(graph[info[1]], info[0])
	}
	for i := 0; i < numCourses && !hasCycle; i++ {
		if visited[i] == false {
			dfs(i)
		}
	}
	// 判断是否可能完成所有课程的学习，如果有环就不能完成
	return !hasCycle
}

func main() {
	pre := [][]int{{1, 0}}
	fmt.Println(canFinish(2, pre))
}

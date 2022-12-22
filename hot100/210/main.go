package main

// 拓扑排序：通常用来“排序”具有依赖关系的任务
func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses) // 邻接表
	// visited 记录哪些节点被遍历过，而 onPath 记录当前递归堆栈中有哪些节点
	visited := make([]bool, numCourses) // 节点是否访问过
	onPath := make([]bool, numCourses)  // 记录当前遍历到的路径
	hasCycle := false                   // 判断是否有环
	var postorder []int
	var dfs func(int)
	dfs = func(s int) {
		// 如果发现 onPath[s] 已经被标记，说明出现了环
		if onPath[s] {
			hasCycle = true
		}
		if visited[s] || hasCycle {
			// 如果已经找到了环，那就不用再遍历了，结束递归
			return
		}
		visited[s] = true
		onPath[s] = true
		for _, t := range graph[s] {
			dfs(t)
		}
		// 后序遍历位置
		postorder = append(postorder, s)
		onPath[s] = false
	}
	// 建立邻接表
	for _, info := range prerequisites {
		graph[info[1]] = append(graph[info[1]], info[0])
	}
	// 开始递归判断是否有环
	for i := 0; i < numCourses; i++ {
		dfs(i)
	}
	if hasCycle {
		return []int{}
	}
	// 反转postorder
	for i, j := 0, len(postorder)-1; i < j; i, j = i+1, j-1 {
		postorder[i], postorder[j] = postorder[j], postorder[i]
	}
	return postorder
}

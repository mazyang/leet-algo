package main

// 岛屿数量
// 淹没法：遇到岛屿计数完之后直接淹没（递归淹没相邻的）
func numIslands(grid [][]byte) int {
	var res int
	m, n := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(row int, col int) {
		// base case 超出索引边界
		if row < 0 || col < 0 || row >= m || col >= n {
			return
		}
		// 如果已经是海水了
		if grid[row][col] == '0' {
			return
		}
		grid[row][col] = '0'
		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row, col-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				res++
				// 找到一个岛屿，将其淹没
				dfs(i, j)
			}
		}
	}
	return res
}

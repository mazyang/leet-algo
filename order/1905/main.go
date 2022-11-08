package main

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
	var res int
	m, n := len(grid2), len(grid2[0])
	var dfs func(int, int)
	dfs = func(row int, col int) {
		// 超出索引边界
		if row < 0 || col < 0 || row >= m || col >= n {
			return
		}
		// 如果已经是海水了
		if grid2[row][col] == 0 {
			return
		}
		grid2[row][col] = 0
		dfs(row+1, col)
		dfs(row-1, col)
		dfs(row, col+1)
		dfs(row, col-1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 如果 grid2某个位置是陆地，但grid1中却不是，说明不是子岛屿，将其淹没
			if grid1[i][j] == 0 && grid2[i][j] == 1 {
				dfs(i, j)
			}
		}
	}
	// 剩下的岛屿就是子岛屿
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid2[i][j] == 1 {
				res++
				dfs(i, j)
			}
		}
	}
	return res
}

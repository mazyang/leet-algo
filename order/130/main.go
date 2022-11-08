package main

import "leetcode-algo/UnionFind"

// 方法一：DFS
func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])
	var dfs func(int, int)
	dfs = func(row int, col int) {
		if col < 0 || col >= n || row < 0 || row >= m || board[row][col] != 'O' {
			return
		}
		board[row][col] = 'F'
		dfs(row-1, col)
		dfs(row+1, col)
		dfs(row, col-1)
		dfs(row, col+1)
	}
	// 遍历边缘一圈
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for i := 0; i < n; i++ {
		dfs(0, i)
		dfs(m-1, i)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// 替换回原来的
			if board[i][j] == 'F' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}

		}
	}
}

func solve2(board [][]byte) {
	m, n := len(board), len(board[0])
	uf := UnionFind.Constructor(m*n + 1)
	dummy := m * n
	// 将首列和末列的0和dummy连通
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' {
			uf.Union(i*n, dummy)
		}
		if board[i][n-1] == 'O' {
			uf.Union(i*n+n-1, dummy)
		}
	}
	// 将首行和末行的0和dummy连通
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' {
			uf.Union(j, dummy)
		}
		if board[m-1][j] == 'O' {
			uf.Union(n*(m-1)+j, dummy)
		}
	}
	// 四个方向
	dir := []struct{ x, y int }{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			// 将此O与上下左右的O连通
			if board[i][j] == 'O' {
				for _, direction := range dir {
					x := i + direction.x
					y := j + direction.y
					if board[x][y] == 'O' {
						uf.Union(x*n+y, i*n+j)
					}
				}
			}
		}
	}
	// 所有不和dummy连通的O,都要被替换
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if !uf.Connected(dummy, i*n+j) {
				board[i][j] = 'X'
			}
		}
	}
}

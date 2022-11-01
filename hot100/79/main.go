package main

// 单词搜索
func exist(board [][]byte, word string) bool {
	var backtrack func(int, int) bool
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	// 指向 word 的索引，当index=len(word)-1 的时候，说明找到了
	index := 0
	backtrack = func(row int, col int) bool {
		// base case 错误的情况，index不可能超过len(word)-1，因为要么出现匹配错误返回false，要么返回true
		if row < 0 || row == m || col < 0 || col == n || visited[row][col] || word[index] != board[row][col] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		// 匹配当前字符成功，继续在四个方向上尝试
		visited[row][col] = true
		index++
		// 其中有一个匹配正确就返回true
		if backtrack(row-1, col) || backtrack(row+1, col) || backtrack(row, col-1) || backtrack(row, col+1) {
			return true
		}
		// 如果四个方向都不匹配，回退一步
		index--
		visited[row][col] = false
		return false
	}
	// 尝试以每个位置作为起点开始回溯
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] != word[0] {
				continue
			}
			if backtrack(i, j) {
				return true
			}
		}
	}
	return false
}

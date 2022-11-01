package main

func solveSudoku(board [][]byte) {
	var backtrack func(int, int) bool
	backtrack = func(row int, col int) bool {
		// 到达最后一列, 换一行继续穷举
		if col == 9 {
			return backtrack(row+1, 0)
		}
		// 行号从0-8，row=9代表结束，找到一个可行解
		if row == 9 {
			return true
		}
		// 如果有预设的数字，换下一个
		if board[row][col] != '.' {
			return backtrack(row, col+1)
		}
		for ch := '1'; ch <= '9'; ch++ {
			// 当前数字不合法，跳过
			if !isValid(board, row, col, byte(ch)) {
				continue
			}
			board[row][col] = byte(ch)
			// 递归寻找可行解
			if backtrack(row, col+1) {
				return true
			}
			board[row][col] = '.'
		}
		return false
	}
	backtrack(0, 0)
}

// 验证放置的位置是否合法
func isValid(board [][]byte, row, col int, num byte) bool {
	// 验证行和列
	for i := 0; i < 9; i++ {
		if board[row][i] == num {
			return false
		}
		if board[i][col] == num {
			return false
		}
		if board[(row/3)*3+i/3][(col/3)*3+i%3] == num {
			return false
		}
	}
	return true
}

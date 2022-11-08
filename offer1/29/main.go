package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	// 代表上下左右边界
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	var res []int
	// 每次遍历一层
	for true {
		// left to right
		for i := l; i < r+1; i++ {
			res = append(res, matrix[t][i])
		}
		t += 1 // 从左到右遍历一层之后，top-1
		if t > b {
			break
		}
		// top to bottom
		for i := t; i < b+1; i++ {
			res = append(res, matrix[i][r])
		}
		r -= 1
		if l > r {
			break
		}
		// right to left
		for i := r; i >= l; i-- {
			res = append(res, matrix[b][i])
		}
		b -= 1
		if t > b {
			break
		}
		//  bottom to top
		for i := b; i >= t; i-- {
			res = append(res, matrix[i][l])
		}
		l += 1
		if l > r {
			break
		}

	}
	return res
}

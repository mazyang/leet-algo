package main

func rotate(matrix [][]int) {
	// 先沿着对角线反转
	n := len(matrix)
	for i := 0; i < n; i++ {
		// j从i开始，代表只遍历对角线上方元素
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// 然后反转每一行
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}

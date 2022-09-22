package main

// 二维区域和检索-矩阵不可变

type NumMatrix struct {
	matrix [][]int
	preSum [][]int // sums 保存以（0,0）为左上角 (row-1, col-1) 为右下角的矩形和
}

func Constructor(matrix [][]int) NumMatrix {
	m, n := len(matrix), len(matrix[0])
	numMatrix := NumMatrix{matrix: matrix}
	preSum := make([][]int, len(matrix)+1)
	for i := 0; i < m+1; i++ {
		preSum[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			preSum[i][j] = preSum[i-1][j] + preSum[i][j-1] - preSum[i-1][j-1] + matrix[i-1][j-1]
		}
	}
	numMatrix.preSum = preSum
	return numMatrix
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.preSum[row2+1][col2+1] - this.preSum[row1][col2+1] - this.preSum[row2+1][col1] + this.preSum[row1][col1]
}

func main() {
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5},
	}
	Constructor(matrix)
}

package main

// 最大矩形，重用 leetcode 84，从上到下计算每一层的最大面积，如果当前最底层等于0，那么高度也是0
func maximalRectangle(matrix [][]byte) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var area int
	m, n := len(matrix), len(matrix[0])
	height := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				height[j] = 0
			} else {
				height[j]++
			}
		}
		area = max(area, largestRectangleArea(height))
	}
	return area
}

// 柱状图的最大矩形
func largestRectangleArea(heights []int) int {
	tmp := make([]int, len(heights)+2)
	copy(tmp[1:], heights)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var area int
	// 单调栈
	var s []int
	for i := 0; i < len(tmp); i++ {
		// 当前元素小于栈顶元素
		for len(s) != 0 && tmp[i] < tmp[s[len(s)-1]] {
			h := tmp[s[len(s)-1]]
			s = s[:len(s)-1]
			left := s[len(s)-1] + 1
			right := i - 1
			area = max(area, (right-left+1)*h)
		}
		s = append(s, i)
	}
	return area
}

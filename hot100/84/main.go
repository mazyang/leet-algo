package main

import "fmt"

// 柱状图的最大矩形
// 方法：单调栈
func largestRectangleArea(heights []int) int {
	// 以每一列作为高度，向左查找第一个小于当前高度的位置，向右查找第一个小于当前高度的位置
	tmp := make([]int, len(heights)+2)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	// 左右补充两个0
	copy(tmp[1:], heights)
	// 单调递增栈，保存元素索引
	var stack []int
	var area int
	for i := 0; i < len(tmp); i++ {
		// 当前元素小于栈顶元素，出栈并计算面积
		for len(stack) != 0 && tmp[i] < tmp[stack[len(stack)-1]] {
			// 当前高度
			h := tmp[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			left := stack[len(stack)-1] + 1
			right := i - 1
			area = max(area, (right-left+1)*h)
		}
		stack = append(stack, i)
	}
	return area
}

func main() {
	height := []int{2, 2, 5, 6, 7, 8}
	fmt.Println(largestRectangleArea(height))
}

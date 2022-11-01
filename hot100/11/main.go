package main

// 盛最多水的容器
// 方法一：暴力
func maxArea(height []int) int {
	// 索引相减获得x轴的值，两边取最小值获得y轴的值
	n := len(height)
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	res := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			res = max((j-i)*min(height[i], height[j]), res)
		}
	}
	return res
}

// 方法二：双指针
func maxArea2(height []int) int {
	left, right := 0, len(height)
	res := 0
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	for left < right {
		// 计算当前窗口
		res = max(res, (right-left)*min(height[left], height[right]))
		// 收缩窗口, 移动较低的一端，因为有可能变大，承的水就可能变多
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

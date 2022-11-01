package main

// 方法一：按列求, 每次求一列中存放的水
// 先获得当前列左边最高的高度，再获得右边最高的高度，某一列存放的水=min(l_max, r_max)-height[i]
func trap(height []int) int {
	n := len(height)
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
	var res int
	// 计算当前列左边最高和右边最高的列(包括当前列)
	lmax := make([]int, n)
	rmax := make([]int, n)
	// base case
	lmax[0] = height[0]
	rmax[n-1] = height[n-1]
	// 从左到右计算（最大高度包括当前高度，便于计算）
	for i := 1; i < n; i++ {
		lmax[i] = max(height[i], lmax[i-1])
	}
	// 从右到左计算
	for i := n - 2; i >= 0; i-- {
		rmax[i] = max(rmax[i+1], height[i])
	}
	// 第一个和最后一个不能存水
	for i := 1; i < n-1; i++ {
		res += min(lmax[i], rmax[i]) - height[i]
	}
	return res
}

// 方法二：栈
func trap2(height []int) int {
	res := 0
	stack := make([]int, 0)
	current := 0
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	for current < len(height) {
		for len(stack) != 0 && height[current] > height[stack[len(stack)-1]] {
			top := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			distance := current - stack[len(stack)-1] - 1
			min := min(height[stack[len(stack)-1]], height[current])
			res += distance * (min - top)
		}
		stack = append(stack, current)
		current++
	}
	return res
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	trap(height)
}

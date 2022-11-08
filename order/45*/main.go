package main

import "fmt"

// 贪心算法
// 跳跃游戏：45 55
func jump(nums []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	end := 0
	farthest := 0
	jumps := 0
	// 在i-end之间选择最优的，
	for i := 0; i < len(nums)-1; i++ {
		// i=0时，farthest=2，end=2，然后在0-2之间寻找最大值赋值给end
		farthest = max(nums[i]+i, farthest)
		if end == i {
			jumps++
			end = farthest
		}
	}
	return jumps
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}

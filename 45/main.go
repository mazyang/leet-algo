package main

import "fmt"

func jump(nums []int) int {
	n := len(nums)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	end := 0
	farthest := 0
	jumps := 0
	for i := 0; i < n-1; i++ {
		// i=2时，farthest=2，end=2，然后在0-2之间寻找最大值赋值给end
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

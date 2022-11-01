package main

func canJump(nums []int) bool {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	farthest := 0
	for i := 0; i < len(nums)-1; i++ {
		// 计算能跳到的最远距离
		farthest = max(farthest, i+nums[i])
		// 可能遇到了0，卡住不动了
		if farthest <= i {
			return false
		}
	}
	return true
}

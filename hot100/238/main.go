package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	// leftSum[i] 代表索引i左侧所有元素的乘积
	leftSum := make([]int, n)
	// rightSum[i] 代表索引i右侧侧所有元素的乘积
	rightSum := make([]int, n)
	leftSum[0] = 1
	for i := 1; i < n; i++ {
		leftSum[i] = leftSum[i-1] * nums[i-1]
	}
	rightSum[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		rightSum[i] = rightSum[i+1] * nums[i+1]
	}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = leftSum[i] * rightSum[i]
	}
	return res
}

func productExceptSelf2(nums []int) []int {
	length := len(nums)
	answer := make([]int, length)
	// answer[i] 表示索引 i 左侧所有元素的乘积
	// 因为索引为 '0' 的元素左侧没有元素， 所以 answer[0] = 1
	answer[0] = 1
	for i := 1; i < length; i++ {
		answer[i] = nums[i-1] * answer[i-1]
	}
	// rightSum 为右侧所有元素的乘积
	// 刚开始右边没有元素，所以 R = 1
	rightSum := 1
	for i := length - 1; i >= 0; i-- {
		// 对于索引 i，左边的乘积为 answer[i]，右边的乘积为 R
		answer[i] = answer[i] * rightSum
		// R 需要包含右边所有的乘积，所以计算下一个结果时需要将当前值乘到 R 上
		rightSum *= nums[i]
	}
	return answer
}

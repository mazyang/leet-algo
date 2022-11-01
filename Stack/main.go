package main

// 单调栈
// 模板：获得下一个更大元素
func nextGreater(nums []int) []int {
	res := make([]int, len(nums))
	var stack []int
	// 单调递减栈
	for i := len(nums) - 1; i >= 0; i-- {
		// 遇到当前元素大于等于栈顶元素，就出栈
		for len(stack) != 0 && nums[i] >= stack[len(stack)-1] {
			// 出栈
			stack = stack[:len(stack)-1]
		}
		// 出栈完成后，栈顶元素就是当前元素的下一个更大元素
		if len(stack) != 0 {
			res[i] = stack[len(stack)-1]
		} else {
			res[i] = -1
		}
		stack = append(stack, nums[i])
	}
	return res
}

// leetcode 496
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	greater := make([]int, len(nums2))
	var stack []int
	// 单调递减栈
	for i := len(nums2) - 1; i >= 0; i-- {
		// 遇到当前元素大于等于栈顶元素，就出栈
		for len(stack) != 0 && nums2[i] >= stack[len(stack)-1] {
			// 出栈
			stack = stack[:len(stack)-1]
		}
		// 出栈完成后，栈顶元素就是当前元素的下一个更大元素
		if len(stack) != 0 {
			greater[i] = stack[len(stack)-1]
		} else {
			greater[i] = -1
		}
		stack = append(stack, nums2[i])
	}
	hashMap := make(map[int]int)
	for i := 0; i < len(nums2); i++ {
		hashMap[nums2[i]] = greater[i]
	}
	res := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		res[i] = hashMap[nums1[i]]
	}
	return res
}

// leetcode 503
func nextGreaterElement2(nums []int) []int {
	res := make([]int, len(nums))
	var stack []int
	n := len(nums)
	// 单调递减栈
	for i := 2*n - 1; i >= 0; i-- {
		// 遇到当前元素大于等于栈顶元素，就出栈
		for len(stack) != 0 && nums[i%n] >= stack[len(stack)-1] {
			// 出栈
			stack = stack[:len(stack)-1]
		}
		// 出栈完成后，栈顶元素就是当前元素的下一个更大元素
		if len(stack) != 0 {
			res[i%n] = stack[len(stack)-1]
		} else {
			res[i%n] = -1
		}
		stack = append(stack, nums[i%n])
	}
	return res
}

// leetcode 739
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	// s存放索引
	var s []int
	for i := n - 1; i >= 0; i-- {
		// 如果当前元素后面的元素小于等于当前元素，出栈
		for len(s) != 0 && temperatures[s[len(s)-1]] <= temperatures[i] {
			s = s[:len(s)-1]
		}
		// 栈为空，不满足条件
		if len(s) == 0 {
			res[i] = 0
		} else {
			// s[len(s)-1] 是比当前元素大的元素的索引，res[i]计算与i之间的距离
			res[i] = s[len(s)-1] - i
		}
		s = append(s, i)
	}
	return res
}

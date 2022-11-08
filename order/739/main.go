package main

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

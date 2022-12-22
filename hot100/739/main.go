package main

func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	var s []int
	for i := n - 1; i >= 0; i-- {
		// 寻找第一个大于当前元素的索引
		if len(s) != 0 && temperatures[i] > temperatures[len(s)-1] {
			s = s[:len(s)-1]
		}
		if len(s) != 0 {
			res[i] = s[len(s)-1] - i
		} else {
			res[i] = 0
		}
		s = append(s, i)
	}
	return res
}

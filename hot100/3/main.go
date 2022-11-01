package main

import "math"

// 使用滑动窗口 寻找最长不重复子串
func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	window := make(map[byte]int)
	res := 0
	for right < len(s) {
		ch1 := s[right]
		right++
		window[ch1] += 1
		// ch1 加入窗口后，窗口中出现重复元素，left 开始改变
		for window[ch1] > 1 {
			ch2 := s[left]
			left++
			window[ch2]--
		}
		res = int(math.Max(float64(res), float64(right-left)))
	}
	return res
}

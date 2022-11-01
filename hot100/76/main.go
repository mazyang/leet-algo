package main

import "math"

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	left, right := 0, 0
	minLen, start := math.MaxInt, 0
	// 代表匹配的元素个数
	match := 0
	for right < len(s) {
		c1 := s[right]
		right++
		// 判断c1 是否在need中，如果在need中，那么就加入window中
		if _, ok := need[c1]; ok {
			window[c1]++
			// 如果c1的个数已经全部匹配，match+1
			if window[c1] == need[c1] {
				match++
			}
		}
		// 是否全部匹配，如果全部匹配，那么开始收缩窗口
		for match == len(need) {
			// 收缩窗口的过程中 更新最小覆盖子串的长度和位置
			if right-left < minLen {
				start = left
				minLen = right - left
			}
			c2 := s[left]
			left++
			// 判断窗口左侧字符是否在need中，如果在，那么收缩窗口的过程中需要更新match的值
			if _, ok := need[c2]; ok {
				if window[c2] == need[c2] {
					match--
				}
				window[c2]--
			}
		}
	}
	if minLen == math.MaxInt {
		return ""
	} else {
		return s[start : start+minLen]
	}
}

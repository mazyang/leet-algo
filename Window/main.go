package main

import "math"

// 滑动窗口相关问题

// 76 最小覆盖子串
func minWindow(s string, t string) string {
	need := map[byte]int{}
	window := map[byte]int{} // window 中的字符必须和need中的相等
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	start := 0
	minLen := math.MaxInt
	left, right := 0, 0
	match := 0 // 窗口中的字符和need中的字符匹配的个数
	for right < len(s) {
		// 扩大窗口
		c1 := s[right]
		right++
		// 判断当前c1是否在need中, 如果need包含c1，那么加入window中
		if _, ok := need[c1]; ok {
			window[c1]++ // 先加1，然后再判断window中的c1是否已经满足need中的c1的个数
			if window[c1] == need[c1] {
				match++
			}
		}
		// window 中的字符已经包含所有的need中的字符, left 一直向前移动直到不满足条件
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

// 567 字符串的排列：判断s2是否包含s1的排列
func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}
	left, right := 0, 0
	match := 0
	for right < len(s2) {
		c1 := s2[right]
		right++
		if _, ok := need[c1]; ok {
			window[c1]++
			if window[c1] == need[c1] {
				match++
			}
		}
		// 窗口大小和s1的长度相等时开始收缩窗口
		for right-left >= len(s1) {
			// 判断窗口中的元素是否和s1匹配
			if match == len(need) {
				return true
			}
			c2 := s2[left]
			left++
			if _, ok := need[c2]; ok {
				if window[c2] == need[c2] {
					match--
				}
				window[c2]--
			}
		}
	}
	return false
}

// 438 找到字符串中所有字母异位词
func findAnagrams(s string, p string) []int {
	var res []int
	need := make(map[byte]int)
	window := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	left, right := 0, 0
	match := 0
	for right < len(s) {
		c1 := s[right]
		right++
		if _, ok := need[c1]; ok {
			window[c1]++
			if window[c1] == need[c1] {
				match++
			}
		}
		// 窗口大小和p的长度相同时，window开始缩小
		for right-left >= len(p) {
			if match == len(need) {
				res = append(res, left)
			}
			c2 := s[left]
			left++
			if _, ok := need[c2]; ok {
				if window[c2] == need[c2] {
					match--
				}
				window[c2]--
			}
		}
	}
	return res
}

// 3 无重复字符的最长子串
func lengthOfLongestSubstring(s string) int {
	left, right := 0, 0
	window := make(map[byte]int)
	res := 0
	for right < len(s) {
		ch1 := s[right]
		right++
		window[ch1] += 1
		// ch1 加入窗口后，窗口中出现重复元素
		for window[ch1] > 1 {
			ch2 := s[left]
			left++
			window[ch2]--
		}
		res = int(math.Max(float64(res), float64(right-left)))
	}
	return res
}

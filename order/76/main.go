package main

import (
	"fmt"
	"math"
)

// 最小覆盖子串
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
			window[c1]++
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

func main() {
	s := "ADOBECODEBANC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
}

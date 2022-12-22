package main

func findAnagrams(s string, p string) []int {
	var res []int
	window := make(map[byte]int)
	need := make(map[byte]int)
	// 用来判断 window 和 need 中的字符匹配情况
	match := 0
	for _, ch := range p {
		need[byte(ch)]++
	}
	left, right := 0, 0
	for right < len(s) {
		ch1 := s[right]
		right++
		if _, ok := need[ch1]; ok {
			window[ch1]++
			if window[ch1] == need[ch1] {
				match++
			}
		}
		if right-left >= len(p) {
			if match == len(need) {
				res = append(res, left)
			}
			ch2 := s[left]
			left++
			if _, ok := need[ch2]; ok {
				if window[ch2] == need[ch2] {
					match--
				}
				window[ch2]--
			}
		}
	}
	return res
}

package main

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

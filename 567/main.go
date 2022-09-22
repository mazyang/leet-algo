package main

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
		for right-left >= len(s1) {
			// 匹配成功
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

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	checkInclusion(s1, s2)
}

package main

func minAddToMakeValid(s string) int {
	leftNeed, rightNeed := 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			rightNeed++
		} else {
			rightNeed--
			// 右括号过多
			if rightNeed == -1 {
				rightNeed = 0
				leftNeed++
			}
		}
	}
	return leftNeed + rightNeed
}

package main

// 有效的括号
func isValid(s string) bool {
	var left []rune
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			left = append(left, c)
		} else {
			// 匹配正确，出栈
			if len(left) != 0 && leftOf(c) == left[len(left)-1] {
				left = left[:len(left)-1]
			} else {
				return false
			}
		}
	}
	return len(left) == 0
}

// 获得右括号对于的左括号
func leftOf(c rune) rune {
	if c == ')' {
		return '('
	} else if c == '}' {
		return '{'
	}
	return '['
}

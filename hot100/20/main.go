package main

// 有效的括号
func isValid(s string) bool {
	var stack []byte
	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, byte(ch))
		} else {
			// 如果左括号栈不为空，并且当前右括号和栈的第一个元素匹配，那就出栈
			if len(stack) != 0 && stack[len(stack)-1] == getAnother(byte(ch)) {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func getAnother(ch byte) byte {
	if ch == ')' {
		return '('
	} else if ch == ']' {
		return '['
	}
	return '{'
}

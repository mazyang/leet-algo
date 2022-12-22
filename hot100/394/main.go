package main

import (
	"strconv"
	"strings"
)

// 字符串解码
func decodeString(s string) string {
	var stk []string
	ptr := 0
	for ptr < len(s) {
		// ptr 指向当前字符
		cur := s[ptr]
		// 如果是数字, 解析一个完整的数字后入栈
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			// 如果当前的字符为字母或者左括号，直接进栈
			stk = append(stk, string(cur))
			ptr++
		} else {
			ptr++
			var sub []string
			// 如果当前字符为右括号，开始出栈直到遇到左括号
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			// 将出栈序列反转后拼接成一个字符串
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			// 删除栈顶的 [
			stk = stk[:len(stk)-1]
			// 取出栈顶的数字, 此时栈顶一定是数字
			repTime, _ := strconv.Atoi(stk[len(stk)-1])
			stk = stk[:len(stk)-1]
			t := strings.Repeat(getString(sub), repTime)
			// 将拼接后的字符串入栈
			stk = append(stk, t)
		}
	}
	return getString(stk)
}

// 读取一个完整的数字
func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

// 拼接字符串
func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}

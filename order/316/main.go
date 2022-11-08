package main

func removeDuplicateLetters(s string) string {
	// 判断当前字符是否已经在单调栈中存在
	inStack := make(map[byte]bool)
	// 记录每个字符的个数, 如果剩余的个数为0，那就不能pop了
	count := make(map[byte]int)
	// 单调栈
	var stack []byte
	for i := 0; i < len(s); i++ {
		count[s[i]]++
	}
	for _, ch := range s {
		// 出现一次，删除一次
		ch := byte(ch)
		count[ch]--
		if _, ok := inStack[ch]; ok && inStack[ch] == true {
			continue
		}
		// 单调递增栈
		for len(stack) != 0 && stack[len(stack)-1] > ch {
			// 当前栈顶元素在之后还会不会出现, 如果不会出现那就不要弹出
			if count[stack[len(stack)-1]] == 0 {
				break
			}
			// 弹出
			inStack[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, ch)
		inStack[ch] = true
	}
	return string(stack)
}

func main() {
	a := "bcabc"
	println(removeDuplicateLetters(a))
}

package main

import (
	"fmt"
	"unicode"
)

func letterCasePermutation(s string) []string {
	var path []byte
	var res []string
	var dfs func(int)
	dfs = func(index int) {
		if index == len(s) {
			res = append(res, string(append([]byte{}, path...)))
			return
		}
		if !unicode.IsLetter(rune(s[index])) {
			path = append(path, s[index])
			dfs(index + 1)
		} else {
			// 做选择
			// 加入小写字母
			path = append(path, byte(unicode.ToLower(rune(s[index]))))
			tmp := len(path) // 因为在后面会加入数字，所以需要保留当前长度
			dfs(index + 1)
			path = path[:tmp-1]

			// 加入大写字母
			path = append(path, byte(unicode.ToUpper(rune(s[index]))))
			tmp = len(path)
			dfs(index + 1)
			path = path[:tmp-1]
		}
	}
	dfs(0)
	return res
}

func main() {
	s := "a1b21"
	fmt.Println(letterCasePermutation(s))
}

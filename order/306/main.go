package main

import (
	"fmt"
	"strconv"
)

func isAdditiveNumber(num string) bool {
	n := len(num)
	var dfs func(a, b int64, num string) bool
	// a b 代表前两个数字
	dfs = func(a, b int64, num string) bool {
		if num == "" {
			return true
		}
		if a+b > 0 && num[0] == '0' {
			return false
		}
		// int64 的最大长度为19
		// 分割的长度最少为1, 所以从1开始
		for i := 1; i < min(len(num)+1, 19); i++ {
			val, _ := strconv.ParseInt(num[:i], 10, 64)
			if a+b == val {
				if dfs(b, val, num[i:]) {
					return true
				}
			}
		}
		return false
	}
	// 以下代码的目的是获得前两个分割的数字，然后再进入回溯
	for i := 1; i < min(n-1, 19); i++ {
		for j := i + 1; j < min(n, i+19); j++ {
			// 如果num[0]=0, 那么第一个数字只能是0, 可以改变的只有第二个数字
			if i > 1 && num[0] == '0' {
				break
			}
			// 如果第二个数字开头为0, continue？为什么不用break
			if j-i > 1 && num[i] == '0' {
				continue
			}
			a, _ := strconv.ParseInt(num[:i], 10, 64)
			b, _ := strconv.ParseInt(num[i:j], 10, 64)
			if dfs(a, b, num[j:]) {
				return true
			}
		}
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	str := "101020"
	fmt.Println(isAdditiveNumber(str))
}

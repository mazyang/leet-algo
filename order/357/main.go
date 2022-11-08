package main

import "fmt"

// 方法一：回溯
func countNumbersWithUniqueDigits(n int) int {
	if n < 0 {
		return 0
	}
	res := 0
	nums := make(map[int]bool)
	for i := 0; i < 10; i++ {
		nums[i] = false
	}
	var count int
	var dfs func(int)
	dfs = func(n int) {
		if count == n {
			res++
			return
		}
		for i := 0; i < 10; i++ {
			// 第一个数字不能为0
			if count == 0 && i == 0 {
				continue
			}
			if nums[i] == false {
				count++
				nums[i] = true
				dfs(n)
				count--
				nums[i] = false
			}
		}
	}
	for i := 0; i <= n; i++ {
		dfs(i)
	}
	return res
}

// 方法二：数学公式，排列组合
func countNumbersWithUniqueDigits2(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	ans, cur := 10, 9
	//如果n=3, 第一位的选择有9种(排除0)，第二位的选择有9种，第三位的选择有8种
	for i := 0; i < n-1; i++ {
		cur *= 9 - i
		ans += cur
	}
	return ans
}

func main() {
	fmt.Println(countNumbersWithUniqueDigits2(2))
}

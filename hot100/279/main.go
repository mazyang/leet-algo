package main

import (
	"fmt"
	"math"
)

// 方法一：回溯
func numSquares(n int) int {
	// 先获得所有小于等于n的完全平方数
	var sqr []int
	for i, tmp := 1, 1; i <= n; i++ {
		tmp = i * i
		if tmp <= n {
			sqr = append(sqr, tmp)
		} else {
			break
		}
	}
	sum, num := 0, 0
	res := math.MaxInt
	var backtrack func(int)
	backtrack = func(start int) {
		if sum > n {
			return
		}
		if sum == n {
			if num < res {
				res = num
			}
			return
		}
		for i := start; i < len(sqr); i++ {
			num++
			sum += sqr[i]
			backtrack(i)
			sum -= sqr[i]
			num--
		}
	}
	backtrack(0)
	return res
}

func numSquares2(n int) int {
	// dp[i] 存放 和为i的完全平方数的最小数量
	dp := make([]int, n+1)
	var sqr []int
	for i, tmp := 1, 1; i <= n; i++ {
		tmp = i * i
		if tmp <= n {
			sqr = append(sqr, tmp)
		} else {
			break
		}
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	for i := 1; i < n+1; i++ {
		dp[i] = n + 1
	}
	for i := 1; i <= n; i++ {
		for _, num := range sqr {
			if i-num < 0 {
				continue
			}
			dp[i] = min(dp[i-num]+1, dp[i])
		}
	}
	return dp[n]
}

// 官方解题
func numSquares3(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		// 在 1 到 i 之间寻找完全平方数
		for j := 1; j*j <= i; j++ {
			minn = min(minn, dp[i-j*j])
		}
		dp[i] = minn + 1
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(numSquares2(12))
}

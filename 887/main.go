package main

import (
	"fmt"
	"math"
)

// 递归+备忘录
func superEggDrop(k int, n int) int {
	memo := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		memo[i] = make([]int, n+1)
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var dp func(int, int) int
	dp = func(K, N int) int {
		if K == 1 {
			return N
		}
		if N == 0 {
			return 0
		}
		res := math.MaxInt
		if memo[K][N] != 0 {
			return memo[K][N]
		}
		for i := 1; i <= N; i++ {
			res = min(res, max(dp(K, N-i), dp(K-1, i-1))+1)
		}
		memo[K][N] = res
		return res
	}
	return dp(k, n)
}

func main() {
	fmt.Println(superEggDrop(2, 10))
}

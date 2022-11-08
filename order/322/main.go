package main

import (
	"fmt"
	"math"
)

// 方法一：暴力
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if amount < 0 {
		return -1
	}
	res := math.MaxInt
	for _, coin := range coins {
		sub := coinChange(coins, amount-coin)
		if sub == -1 {
			continue
		}
		res = int(math.Min(float64(res), float64(sub)+1))
	}
	if res == math.MaxInt {
		return -1
	} else {
		return res
	}
}

// 方法二：递归，自顶向下+备忘录
func coinChange2(coins []int, amount int) int {
	memo := make([]int, amount+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = -100
	}
	var dfs func(int) int
	dfs = func(amount int) int {
		if amount == 0 {
			return 0
		}
		if amount < 0 {
			return -1
		}
		if memo[amount] != -100 {
			return memo[amount]
		}
		res := math.MaxInt
		for _, coin := range coins {
			sub := dfs(amount - coin)
			if sub == -1 {
				continue
			}
			res = int(math.Min(float64(res), float64(sub)+1))
		}
		if res == math.MaxInt {
			return -1
		} else {
			return res
		}
	}
	return dfs(amount)
}

// 方法三：迭代，自底向上
func coinChange3(coins []int, amount int) int {
	// dp[i] 代表总额为i时所需要的最少硬币
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	min := func(x, y int) int {
		if x > y {
			return y
		}
		return x
	}
	dp[0] = 0
	for i := 0; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], 1+dp[i-coin])
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}
func main() {
	fmt.Println(coinChange3([]int{1, 2, 5}, 100))
}

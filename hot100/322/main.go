package main

func coinChange(coins []int, amount int) int {
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	// dp[i]代表总金额为i时需要的零钱数量
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] == amount+1 {
		return -1
	} else {
		return dp[amount]
	}
}

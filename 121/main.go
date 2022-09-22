package main

import "math"

// 原始版本
func maxProfit(prices []int) int {
	// dp[i][j] 代表第i天的时候持有（j=1）或 未持有股票（j=0）时的总收益
	dp := make([][]int, len(prices))
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 2)
	}
	for i := 0; i < len(prices); i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[0]
			continue
		}
		// 如果第i天未持有，第i-1天可能未持有，也可能持有
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 如果第i天持有
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[len(prices)-1][0]
}

// 状态压缩
func maxProfit2(prices []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	// base case：dp[-1][0]=0  dp[-1][1]=负无穷
	dp0 := 0
	dp1 := math.MinInt
	for i := 0; i < len(prices); i++ {
		// 如果第i天未持有，第i-1天可能未持有，也可能持有
		dp0 = max(dp0, dp1+prices[i])
		// 如果第i天持有
		dp1 = max(dp1, -prices[i])
	}
	return dp0
}

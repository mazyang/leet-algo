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
		if i == 1 {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = max(dp[i-1][1], -prices[i])
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 冷冻期：卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)
		// 如果当前持有，前一天如果没有持有，有可能是刚卖出
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
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
	dp0pre := 0
	for i := 0; i < len(prices); i++ {
		tmp := dp0
		dp0 = max(dp0, dp1+prices[i])
		dp1 = max(dp1, dp0pre-prices[i])
		dp0pre = tmp
	}
	return dp0
}

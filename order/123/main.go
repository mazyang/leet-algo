package main

import "math"

// 原始版本
func maxProfit(prices []int) int {
	dp := make([][][]int, len(prices))
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := 0; i < len(prices); i++ {
		dp[i] = make([][]int, 3)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < len(prices); i++ {
		for k := 1; k <= 2; k++ {
			if i == 0 {
				dp[i][k][0] = 0
				dp[i][k][1] = -prices[0]
				continue
			}
			dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1]+prices[i])
			dp[i][k][1] = max(dp[i-1][k][1], dp[i-1][k-1][0]-prices[i])
		}
	}
	return dp[len(prices)-1][2][0]
}

// 状态压缩
func maxProfit2(prices []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	// base case
	dp_i10 := 0
	dp_i11 := math.MinInt
	dp_i20 := 0
	dp_i21 := math.MinInt
	for i := 0; i < len(prices); i++ {
		dp_i20 = max(dp_i20, dp_i21+prices[i])
		dp_i21 = max(dp_i21, dp_i10-prices[i])
		dp_i10 = max(dp_i10, dp_i11+prices[i])
		dp_i11 = max(dp_i11, -prices[i])
	}
	return dp_i20
}

package main

// 买卖股票的最佳时机含手续费
// 买入的时候支付手续费
func maxProfit(prices []int, fee int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	n := len(prices)
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i] - fee
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i]-fee)
	}
	return dp[n-1][0]
}

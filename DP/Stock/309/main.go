package main

// 最佳买卖股票时机含冷冻期
func maxProfit(prices []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	n := len(prices)
	dp := make([][2]int, n)
	for i := 1; i < n; i++ {
		if i-1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		if i-2 == -1 {
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = max(dp[i-1][1], -prices[i])
		}
		// 当前没有持有股票，1.昨天也没有持有  2.昨天就持有，今天卖出
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 当前持有股票，1.昨天就持有  2.前天没有持有(可能刚卖出，隔一天才能买入)，今天买入
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
	}
	return dp[n-1][0]
}

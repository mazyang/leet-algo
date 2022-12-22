package main

// 买卖股票的最佳时机，k=正无穷
// 动态规划
func maxProfit(prices []int) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	n := len(prices)
	dp := make([][2]int, n)
	for i := 0; i < n; i++ {
		if i-1 == -1 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}
		// 当前没有持有股票，1.昨天也没有持有  2.昨天就持有，今天卖出
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		// 当前持有股票，1.昨天就持有  2.昨天没有持有，今天买入
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}

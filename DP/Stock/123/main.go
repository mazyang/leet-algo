package main

// 买卖股票的最佳时机，k=2
func maxProfit(prices []int) int {
	k := 2
	n := len(prices)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	dp := make([][][2]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][2]int, k+1)
	}
	for i := 0; i < n; i++ {
		for j := k; j >= 1; j-- {
			if i-1 == -1 {
				dp[i][j][0] = 0
				dp[i][j][1] = -prices[i]
				continue
			}
			// 当前没有持有股票，1.昨天也没有持有  2.昨天就持有，今天卖出，所以+prices[i]
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			// 当前持有股票，1.昨天就持有  2.昨天没有持有，今天买入(买入的时候交易次数减1)，所有-prices[i]
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[n-1][k][0]
}

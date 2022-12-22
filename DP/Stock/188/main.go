package main

// 买卖股票的最佳时机，k=任意值
// 每次股票买入时交易次数减1
func maxProfit(k int, prices []int) int {
	n := len(prices)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	// dp[2][3][0]代表在第2天的时候，至今最多进行3次交易，没有持有股票 的最大收益
	// 每次股票买入时交易次数减1
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
			// 当前持有股票，1.昨天就持有  2.昨天没有持有，今天买入，所以交易次数+1
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}
	return dp[n-1][k][0]
}

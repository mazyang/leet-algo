package main

// 爬楼梯
func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	// dp[i] 代表有i个楼梯时需要的方法个数
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	// 在第i层的时候，可能是在i-1或i-2层跳上来的
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

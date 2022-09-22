package main

// 方法一：自顶向下
func fib(n int) int {
	memo := make([]int, n+1)
	var dfs func(int) int
	dfs = func(i int) int {
		if i == 0 || i == 1 {
			return i
		}
		if memo[i] != 0 {
			return memo[i]
		}
		memo[i] = dfs(i-1) + dfs(i-2)
		return memo[i]
	}
	return dfs(n)
}

// 方法二：自底向上
func fib2(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	dp1, dp2 := 0, 1
	for i := 2; i <= n; i++ {
		dp1, dp2 = dp2, dp1+dp2
	}
	return dp2
}

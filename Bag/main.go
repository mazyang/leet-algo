package main

import "fmt"

// 背包可以承受的最大重量W, 给N个物品，wt val 分别是N个物体的重量和价值
// 求背包最多能装的价值是多少?
func knapsack(W int, N int, wt []int, val []int) int {
	// dp[i][w] 的定义如下：对于前 i (0...i-1)个物品，当前背包的容量为 w，这种情况下可以装的最⼤价值是 dp[i][w]
	dp := make([][]int, N+1)
	// dp[0][..]=dp[..][0]=0
	for i := 0; i < N+1; i++ {
		dp[i] = make([]int, W+1)
	}
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	// 对于当前物品 i-1，可以装入或不装入
	for i := 1; i <= N; i++ {
		// 背包容量从1开始计算
		for w := 1; w <= W; w++ {
			// 当前背包的容量小于当前物品的重量，不装入
			if w-wt[i-1] < 0 {
				dp[i][w] = dp[i-1][w]
			} else {
				// 装入或者不装入，选择最大的
				dp[i][w] = max(dp[i-1][w-wt[i-1]]+val[i-1], dp[i-1][w])
			}
		}
	}
	return dp[N][W]
}

func main() {
	N := 3
	W := 4
	wt := []int{2, 1, 3}
	val := []int{4, 2, 3}
	fmt.Println(knapsack(W, N, wt, val))
}

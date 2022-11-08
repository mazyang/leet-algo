package main

import "fmt"

// 不同的二叉搜索树
// 方法一：暴力递归，以从1到n其中一个数字作为根节点递归获得子树
func numTrees(n int) int {
	var count func(int, int) int
	count = func(low int, high int) int {
		if low > high {
			return 1
		}
		res := 0
		for i := low; i <= high; i++ {
			left := count(low, i-1)
			right := count(i+1, high)
			res += left * right
		}
		return res
	}
	return count(1, n)
}

// 方法二：暴力+备忘录
func numTrees1(n int) int {
	memo := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = make([]int, n+1)
	}
	// 备忘录的值初始化为 0，备忘录中存放从 lo到hi 构成的二叉搜索树的个数
	var count func(int, int) int
	count = func(lo int, hi int) int {
		if lo > hi {
			return 1
		}
		// 查备忘录
		if memo[lo][hi] != 0 {
			return memo[lo][hi]
		}
		res := 0
		for mid := lo; mid <= hi; mid++ {
			left := count(lo, mid-1)
			right := count(mid+1, hi)
			res += left * right
		}
		// 将结果存入备忘录
		memo[lo][hi] = res
		return res
	}
	return count(1, n)
}

// 动态规划
func numTrees2(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}
func main() {
	fmt.Println(numTrees(6))
}

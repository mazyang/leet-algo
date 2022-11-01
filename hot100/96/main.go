package main

import "fmt"

// 方法一：递归
func numTrees(n int) int {
	// base case
	if n <= 1 {
		return 1
	}
	num := 0
	for i := 1; i <= n; i++ {
		// 以i作为根节点分割
		leftNum := numTrees(i - 1)
		rightNum := numTrees(n - i)
		num += leftNum * rightNum
	}
	return num
}

// 方法二：动态规划
func numTrees2(n int) int {
	// dp[i] 记录从1到i互不相同的二叉搜索树的个数
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		// 以j作为分割点，j-1 为左子树的节点个数，i-j 为右子树节点的个数
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

// 方法三：备忘录+递归
func numTrees3(n int) int {
	memo := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		memo[i] = make([]int, n+1)
	}
	// 备忘录的值初始化为 0
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

func main() {
	fmt.Println(numTrees2(4))
}

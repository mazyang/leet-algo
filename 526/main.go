package main

import (
	"fmt"
	"math/bits"
)

func countArrangement(n int) int {
	// 构造全排列数组，构造的同时判断是否满足题意
	var res int
	var dfs func(int)
	visited := make([]bool, n+1)
	dfs = func(index int) {
		if index > n {
			res++
			return
		}
		// i代表位置，index代表数字
		for i := 1; i <= n; i++ {
			if visited[i] == false && (i%index == 0 || index%i == 0) {
				visited[i] = true
				dfs(index + 1)
				visited[i] = false
			}
		}
	}
	dfs(1)
	return res
}

func countArrangement2(n int) int {
	ans := 0
	match := make(map[int][]int)
	// i代表位置，j代表所在位置满足条件的数字
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i%j == 0 || j%i == 0 {
				match[i] = append(match[i], j)
			}
		}
	}
	vis := make([]bool, n+1)

	var dfs func(i int)
	// i代表位置
	dfs = func(i int) {
		if i == n+1 {
			ans++
			return
		}
		// 向位置i中放入j
		for _, j := range match[i] {
			if !vis[j] {
				vis[j] = true
				dfs(i + 1)
				vis[j] = false
			}
		}
	}

	dfs(1)
	return ans
}

// 方法三：动态规划
func countArrangement3(n int) int {
	f := make([]int, 1<<n)
	f[0] = 1
	for mask := 1; mask < 1<<n; mask++ {
		num := bits.OnesCount(uint(mask))
		for i := 0; i < n; i++ {
			if mask>>i&1 > 0 && (num%(i+1) == 0 || (i+1)%num == 0) {
				f[mask] += f[mask^1<<i]
			}
		}
	}
	return f[1<<n-1]
}

func main() {
	fmt.Println(countArrangement(4))
}

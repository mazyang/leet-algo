package main

import (
	"fmt"
	"sort"
)

func lengthOfLIS(nums []int) int {
	n := len(nums)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	res := 1
	// dp[i]代表以nums[i]结尾的一个最长递增子序列的长度
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}

// 方法一：使用 LIS
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		if a[0] == b[0] {
			return a[1] > b[1]
		}
		return a[0] < b[0]
	})
	var height []int
	for i := 0; i < len(envelopes); i++ {
		height = append(height, envelopes[i][1])
	}
	return lengthOfLIS(height)
}

// 方法二：二分查找+动态
func maxEnvelopes2(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		a, b := envelopes[i], envelopes[j]
		return a[0] < b[0] || a[0] == b[0] && a[1] > b[1]
	})

	var f []int
	for _, e := range envelopes {
		// e[1] = 3 4 7 4
		h := e[1]
		// 如果h在f中没有查找到，那么返回的i就是在f中待插入的位置
		if i := sort.SearchInts(f, h); i < len(f) {
			f[i] = h
		} else {
			f = append(f, h)
		}
	}
	return len(f)
}

func main() {
	envelopes := [][]int{{5, 2}, {6, 7}, {6, 4}, {1, 8}, {2, 3}, {5, 4}}
	fmt.Println(maxEnvelopes2(envelopes))
}

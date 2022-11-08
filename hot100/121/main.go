package main

import "math"

// 股票买卖的最佳时机
func maxProfit(prices []int) int {
	res := 0
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	minPrice := math.MaxInt
	// 遍历的时候计算最低点，然后每次计算利润
	for i := 0; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		res = max(res, prices[i]-minPrice)
	}
	return res
}

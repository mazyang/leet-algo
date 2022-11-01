package main

import "sort"

// 区间调度相关问题
// 给你很多形如[start,end]的闭区间，请你设计一个算法，算出这些区间中最多有几个互不相交的区间

func intervalScheduling(intervals [][]int) int {
	// 1.先按照end升序排列
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	count := 0
	// 从区间集合 intervals 中选择一个区间 x，这个 x 是在当前所有区间中结束最早的（end 最小）
	x_end := intervals[0][1]
	// 2.遍历区间开始寻找互不相交的区间
	for _, interval := range intervals {
		start := interval[0]
		// 新的区间起始位置大于上一个区间的结束位置的时候，代表两个区间是互不相交的区间
		if start >= x_end {
			count++
			x_end = interval[1]
		}
	}
	return count
}

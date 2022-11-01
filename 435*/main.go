package main

import "sort"

// 计算互不相交的区间个数
func eraseOverlapIntervals(intervals [][]int) int {
	// 按照 end 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	end := intervals[0][1]
	// 计算不相交的区间个数
	count := 1
	for _, interval := range intervals {
		start := interval[0]
		// 代表不相交
		if start >= end {
			count += 1
			end = interval[1]
		}
	}
	// 需要移除的区间个数
	return len(intervals) - count
}

package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	// 按照 end 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	end := intervals[0][1]
	// 不相交的个数
	count := 1
	for _, interval := range intervals {
		start := interval[0]
		if start >= end {
			count += 1
			end = interval[1]
		}
	}
	return len(intervals) - count
}

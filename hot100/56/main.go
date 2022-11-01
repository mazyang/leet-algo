package main

import "sort"

// 合并区间
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[i][0]
	})
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var res [][]int
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		last := res[len(res)-1]
		if cur[0] <= last[1] {
			last[1] = max(cur[1], last[1])
		} else {
			res = append(res, cur)
		}
	}
	return res
}

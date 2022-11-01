package main

import "sort"

// 区间问题，就是线段问题，让你合并所有线段、找出线段的交集等等

// 1288 删除被覆盖区间
func removeCoveredIntervals(intervals [][]int) int {
	// 按照起点升序排列，如果起点相同，则按照终点降序排列
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] > intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	left, right := intervals[0][0], intervals[0][1]
	res := 0
	// 从第二个区间开始计算
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		// 情况一：找到被覆盖区间
		if left <= interval[0] && right >= interval[1] {
			res++
		}
		// 情况二：找到相交区间，合并
		if right >= interval[0] && right <= interval[1] {
			right = interval[1]
		}
		// 情况三：区间完全不相交，更新起点和终点
		if right <= interval[0] {
			left = interval[0]
			right = interval[1]
		}
	}
	return len(intervals) - res
}

// 56 合并区间
func merge(intervals [][]int) [][]int {
	// 1.先按照start升序排列
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var res [][]int
	res = append(res, intervals[0])
	// 2.遍历合并区间
	for i := 1; i < len(intervals); i++ {
		// 读取到的当前区间
		cur := intervals[i]
		// 从res中获得最后一个区间
		last := res[len(res)-1]
		if cur[0] <= last[1] {
			// 寻找最大的end
			last[1] = max(cur[1], last[1])
		} else {
			res = append(res, cur)
		}
	}
	return res
}

// 986 区间列表的交集
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	i, j := 0, 0
	var res [][]int
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
	// 使用两个指针
	for i < len(firstList) && j < len(secondList) {
		a1, a2 := firstList[i][0], firstList[i][1]
		b1, b2 := secondList[j][0], secondList[j][1]
		// 存在交集，使用 max min 获得交集
		if b2 >= a1 && a2 >= b1 {
			res = append(res, []int{max(a1, b1), min(a2, b2)})
		}
		if b2 < a2 {
			j++
		} else {
			i++
		}
	}
	return res
}

package main

import "sort"

func videoStitching(clips [][]int, time int) int {
	// 按起点升序排列，起点相同的降序排列
	sort.Slice(clips, func(i, j int) bool {
		if clips[i][0] == clips[j][0] {
			return clips[i][1] > clips[j][1]
		}
		return clips[i][0] < clips[j][0]
	})
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	var res int
	curEnd, nextEnd := 0, 0
	for i := 0; i < len(clips) && clips[i][0] <= curEnd; {
		for i < len(clips) && clips[i][0] <= curEnd {
			nextEnd = max(nextEnd, clips[i][1])
			i++
		}
		res++
		curEnd = nextEnd
		if curEnd >= time {
			return res
		}
	}
	return -1
}

func main() {
	clips := [][]int{{0, 2}, {4, 6}, {8, 10}, {1, 9}, {1, 5}, {5, 9}}
	time := 10
	videoStitching(clips, time)
}

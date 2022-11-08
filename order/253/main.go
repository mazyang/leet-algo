package main

import (
	"math"
	"sort"
)

func minMeetingRooms(meetings [][]int) int {
	n := len(meetings)
	begin := make([]int, n)
	end := make([]int, n)
	for i := 0; i < n; i++ {
		begin[i] = meetings[i][0]
		end[i] = meetings[i][1]
	}
	sort.Ints(begin)
	sort.Ints(end)
	count, res := 0, 0
	for i, j := 0, 0; i < n && j < n; {
		if begin[i] < end[j] {
			count++
			i++
		} else {
			count--
			j++
		}
		res = int(math.Max(float64(res), float64(count)))
	}
	return res
}

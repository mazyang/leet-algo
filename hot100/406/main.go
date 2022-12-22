package main

import "sort"

func reconstructQueue(people [][]int) [][]int {
	// 按照身高降序排列
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})
	// ans 是已经有序的集合
	var ans [][]int
	for _, person := range people {
		// 在ans的idx位置插入一个元素
		idx := person[1]
		//tmp := append([][]int{person}, ans[idx:]...)
		//ans = append(ans[:idx], tmp...)
		ans = append(ans, []int{})
		copy(ans[idx+1:], ans[idx:])
		ans[idx] = person
	}
	return ans
}

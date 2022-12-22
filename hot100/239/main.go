package main

import "container/list"

// 单调递减队列，新加入的元素要大于队尾元素
func push(l *list.List, n int) {
	for l.Len() != 0 && n > l.Back().Value.(int) {
		tail := l.Back()
		l.Remove(tail)
	}
	l.PushBack(n)
}

func pop(l *list.List, n int) {
	if n == l.Front().Value.(int) {
		front := l.Front()
		l.Remove(front)
	}
}

func max(l *list.List) int {
	return l.Front().Value.(int)
}

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	window := list.New()
	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			push(window, nums[i])
		} else {
			push(window, nums[i])
			res = append(res, max(window))
			pop(window, nums[i-k+1])
		}
	}
	return res
}

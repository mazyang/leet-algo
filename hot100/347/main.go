package main

import "container/heap"

type IHeap [][2]int

func (h IHeap) Len() int {
	return len(h)
}

// Less 比较频率
func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	hashMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		hashMap[nums[i]]++
	}
	h := &IHeap{}
	heap.Init(h)
	for key, value := range hashMap {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	var res []int
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(h).([2]int)[0])
	}
	return res
}

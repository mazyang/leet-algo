package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"sort"
)

type IntHeap sort.IntSlice

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	// *h 代表指针h指向的对象
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 方法一：最小堆维持k个元素，遍历完成后堆顶就是第K大的元素
func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range nums {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h)[0]
}

// 方法二：快速排序
func findKthLargest2(nums []int, k int) int {
	n := len(nums)
	target := n - k // target 是第k大的元素的下标
	left, right := 0, n-1
	for {
		pivotIdx := partition(nums, left, right)
		// 找到目标元素
		if pivotIdx == target {
			return nums[pivotIdx]
		} else if pivotIdx < target {
			left = pivotIdx + 1
		} else {
			right = pivotIdx - 1
		}
	}
}

// 返回分界点的索引
func partition(nums []int, left, right int) int {
	// 也可以先洗牌
	randomIdx := left + rand.Intn(right-left+1)
	// 交换第一个元素和随机位置的元素
	nums[left], nums[randomIdx] = nums[randomIdx], nums[left]
	// 获得分界点
	pivot := nums[left]
	i, j := left+1, right
	for i <= j {
		// 从左到右寻找第一个大于分界点的元素
		for i <= j && nums[i] <= pivot {
			i++
		}
		// 从右到左寻找第一个小于分界点的元素
		for i <= j && nums[j] > pivot {
			j--
		}
		if i >= j {
			break
		}
		// 找到之后交换，继续寻找
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[left], nums[j] = nums[j], nums[left]
	return j
}

func main() {
	//nums := []int{3, 2, 1, 5, 6, 4}
	//fmt.Println(findKthLargest(nums, 2))
	h := &IntHeap{}
	heap.Init(h)
	heap.Push(h, 3)
	heap.Push(h, 2)
	heap.Push(h, 1)
	heap.Push(h, 5)
	heap.Push(h, 6)
	heap.Push(h, 4)
	fmt.Println(heap.Pop(h))
}

package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

// NodeHeap 是一个由ListNode组成的最小堆
type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 合并K个升序链表
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	h := &NodeHeap{}
	heap.Init(h)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}
	res := &ListNode{Val: -1}
	p := res
	for h.Len() != 0 {
		// 从最小堆中弹出
		node := heap.Pop(h)
		p.Next = node.(*ListNode)
		if node.(*ListNode).Next != nil {
			heap.Push(h, node.(*ListNode).Next)
		}
		p = p.Next
	}
	return res.Next
}

// 这个示例会将一些整数插入到堆里面， 接着检查堆中的最小值，
// 之后按顺序从堆里面移除各个整数。
func main() {

}

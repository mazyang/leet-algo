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
	// *h 代表指针h指向的对象
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 合并K个升序链表
// 方法：优先队列，最小堆
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	h := &NodeHeap{}
	heap.Init(h)
	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}
	dummy := &ListNode{Val: -1}
	p := dummy
	for h.Len() != 0 {
		// 弹出一个节点
		node := heap.Pop(h).(*ListNode)
		// 加入新的链表中
		p.Next = node
		// 往堆中加入新节点
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
		p = p.Next
	}
	return dummy.Next
}
